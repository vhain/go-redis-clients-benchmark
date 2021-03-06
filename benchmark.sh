#!/bin/bash

set -euxo pipefail

TMPFILE="$(mktemp)"
go test -bench=. -benchtime=10000x | tee "${TMPFILE}"
echo "# Benchmark Result" > BENCHMARK.md
echo "Benchmark result generated by \`make bench\`" >> BENCHMARK.md
echo "\`\`\`" >> BENCHMARK.md
cat "${TMPFILE}" >> BENCHMARK.md
echo "\`\`\`" >> BENCHMARK.md
