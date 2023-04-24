.PHONY: bench
bench: clean BENCHMARK.md

clean:
	rm -f benchmark.txt BENCHMARK.md

benchmark.txt:
	go test -bench=. -benchtime=10000x | tee $@

BENCHMARK.md: benchmark.txt
	echo "# Benchmark Result" > $@
	echo "Benchmark result generated by \`make bench\`" >> $@
	echo "\`\`\`" >> $@
	cat benchmark.txt >> $@
	echo "\`\`\`" >> $@
