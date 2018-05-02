#!/usr/bin/env bash

########################
# include the magic
########################
. demo-magic.sh


########################
# Configure the options
########################

#
# speed at which to simulate typing. bigger num = faster
#
TYPE_SPEED=80

#
# custom prompt
#
# see http://www.tldp.org/HOWTO/Bash-Prompt-HOWTO/bash-prompt-escape-sequences.html for escape sequences
#
DEMO_PROMPT="${GREEN}➜ ${CYAN}\W "

# hide the evidence
clear

# Stage 1: Quine
pe "cd stage1"
pe "go run quine.go"
pe "go run quine.go > newquine.go"
pe "diff quine.go newquine.go"
pe "cd .."

# Stage 2: Compiler bootstrapping
pe "cd stage2"
pe "go build -o compiler ../compiler.go"
pe "./compiler build -o hw hw.go"
pe "./hw"
pe "./compiler build -o hw-fetch hw-fetch.go"

pe "./compiler build -o compiler training-compiler.go"
pe "./compiler build -o hw-fetch hw-fetch.go"
pe "./hw-fetch"

pe "./compiler build -o compiler trained-compiler.go"
pe "./compiler build -o hw-fetch hw-fetch.go"
pe "./hw-fetch"
pe "cd .."

# Stage 3: Conduct the attack
pe "cd stage3"
pe "go build -o login login.go"
pe "./login monkey"
pe "./login backdoor"
pe "go build -o login login-hacked.go"
pe "./login backdoor"

pe "go build -o compiler compiler-hack-login.go"
pe "./compiler build -o login login.go"
pe "./login backdoor"

pe "go build -o compiler compiler-hack-itself.go"
pe "./compiler build -o login login.go"
pe "./login backdoor"
pe "./compiler build -o compiler ../compiler.go"
pe "./compiler build -o login login.go"
pe "./login backdoor"

# Stage 4: Subverting verification
pe "shasum -a 256 /usr/local/bin/go"
pe "shasum -a 256 compiler"

pe "cd .."
pe "cd stage4"
pe "go build -o mysha256 mysha256.go"
pe "./mysha256 /usr/local/bin/go"
pe "./mysha256 ../stage3/compiler"

pe "go build -o mysha256 mysha256-hacked.go"
pe "./mysha256 /usr/local/bin/go"
pe "./mysha256 ../stage3/compiler"

pe "go build -o compiler compiler-hack-ultimate.go"
pe "./compiler build -o mysha256 mysha256.go"
pe "./mysha256 compiler"
pe "shasum -a 256 compiler"
pe "./compiler build -o compiler ../compiler.go"
pe "./compiler build -o compiler ../compiler.go"
pe "./compiler build -o mysha256 mysha256.go"
pe "./mysha256 compiler"
