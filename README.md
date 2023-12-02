# rproxy

# Testing the dummy reverse proxy in the following steps:
Step 1: on one terminal, do "git clone https://github.com/michael4RD/rproxy.git"

Step 2: stay on the same terminal, and do "cd rproxy && make run"

Step 3: open another terminal and type "curl localhost:8080 -v"

Note if you hit errors like "listen tcp :8080: bind: address already in use",
please do "sudo lsof -i :8080" or "sudo lsof -i : 8888", followed by "sudo kill [pid]",
and redo step 2&3.

# Resources: how was the implementation built upon?
Researched into the topic with online resources, and applied the "trial-and-fail" approach.

# Design decisions and limitations of the system
The goal is to build an in-house HTTP reverse proxy, without leveraging external packages.
I decide to build the minimum valuable product version of the HTTP reverse proxy: dummy case,
to start with. We can extend the design to support more features (e.g., throttling, load balancing,
authentication, routing policy enforcement, etc) in the future, based on the dummy one.

The current implementation has several drawbacks: it does not scale, not secure, and does not
support advanced features yet.

# How to make it scale?
We can make the RProxy stateless to make it scalable: multiple stateless RProxy instances
can run in parallel (preferrably in a docker container environment) without maintaining
state consistency in-between.

We can also make the currently single target server scalable by supporting a list of target
URLs. We can redirect incoming requests to different target servers based on rules.

# How to make it secure?
We can change the code to use "http.ListenAndServeTLS" instead of "http.ListenAndServe" to
add transport layer security to the HTTP requests (i.e., https://localhost:8080/8088).
