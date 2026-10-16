#!/bin/bash
# Downloads the Julep OpenAPI spec into api/. The spec is not committed;
# `go generate` fetches it as its first step. Stainless hosts the spec under a
# URL that carries a content hash and changes on every update, so the current
# URL is read from the .stats.yml of Julep's Python SDK first.
set -euo pipefail
trap 'on_error "$?" "$LINENO"' ERR

LOG_FILE="${LOG_FILE:-/tmp/$(basename "$0" .sh).log}"
exec > >(tee -a "$LOG_FILE") 2>&1

readonly STATS_URL="https://raw.githubusercontent.com/julep-ai/python-sdk/main/.stats.yml"
readonly DOWNLOAD_TIMEOUT_SECONDS=60
readonly SPEC_NAME="openapi.yaml"
SPEC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/api"
readonly SPEC_DIR

on_error() {
	local exit_code="$1"
	local line="$2"

	printf '{"level":"ERROR","file":"%s","line":%d,"msg":"command failed","exit_code":%d}\n' \
		"${BASH_SOURCE[1]}" "$line" "$exit_code" >&2
}

log() {
	local level="$1"
	shift

	jq -nc \
		--arg time "$(date -u '+%Y-%m-%dT%H:%M:%SZ')" \
		--arg level "$level" \
		--arg file "${BASH_SOURCE[1]##*/}" \
		--arg line "${BASH_LINENO[0]}" \
		--arg func "${FUNCNAME[1]:-main}" \
		--arg msg "$*" \
		'{time:$time,level:$level,file:$file,line:($line|tonumber),func:$func,msg:$msg}' >&2
}

spec_url() {
	local stats url
	stats="$(wget -q -T "$DOWNLOAD_TIMEOUT_SECONDS" -O - "$STATS_URL")"
	url="$(printf '%s\n' "$stats" | sed -n 's/^openapi_spec_url:[[:space:]]*//p' | tr -d "\"'\r")"

	if [[ "$url" != https://* ]]; then
		log ERROR "no openapi_spec_url in stats file url=${STATS_URL}"
		exit 1
	fi

	printf '%s\n' "$url"
}

main() {
	mkdir -p "$SPEC_DIR"
	local tmp url
	tmp="$(mktemp "${SPEC_DIR}/.${SPEC_NAME}.XXXXXX")"
	# shellcheck disable=SC2064 # expand tmp now, the local is gone when EXIT fires
	trap "rm -f -- '${tmp}'" EXIT

	url="$(spec_url)"
	log INFO "downloading spec=${SPEC_NAME} url=${url}"
	wget -q \
		-T "$DOWNLOAD_TIMEOUT_SECONDS" \
		-O "$tmp" \
		"$url"

	# The spec is YAML written in JSON style and the image has no YAML parser,
	# so check for the openapi key instead of parsing it.
	if ! grep -qE '^[[:space:]]*"?openapi"?:' "$tmp"; then
		log ERROR "downloaded file is not an OpenAPI document spec=${SPEC_NAME}"
		exit 1
	fi

	chmod 0644 "$tmp"
	mv -- "$tmp" "${SPEC_DIR}/${SPEC_NAME}"
	log INFO "saved spec=${SPEC_NAME} openapi=$(grep -m1 -oE '"?openapi"?: *"?[0-9.]+' "${SPEC_DIR}/${SPEC_NAME}" | grep -oE '[0-9.]+$')"
}

main "$@"
