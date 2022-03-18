run_alice:
	go run ./alice/alice.go

run_bob:
	go run ./bob/bob.go

setup:
	chmod +x ./generate_keys.sh
	./generate_keys.sh