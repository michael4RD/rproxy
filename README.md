# rproxy

# Goal: Build an in-house HTTP reverse proxy, without leveraging external packages

# Test: Testing the dummy reverse proxy in the following steps:
Step 1: on one terminal, do "git clone https://github.com/michael4RD/rproxy.git"

Step 2: stay on the same terminal, and do "cd rproxy && make run"

Step 3: open another terminal and type "curl localhost:8080 -v"

Note if you hit errors like "listen tcp :8080: bind: address already in use",
please do "sudo lsof -i :8080" or "sudo lsof -i : 8888", following by "sudo kill <pid>"

#
