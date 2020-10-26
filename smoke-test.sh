#!/bin/bash
curl -L https://github.com/ThoughtWorks-DPS/poc-va-cli/releases/download/v0.1.29/poc-va-cli_0.1.29_Linux_i386.tar.gz --output poc-va-cli_0.1.29_Linux_i386.tar.gz
tar -xvf poc-va-cli_0.1.29_Linux_i386.tar.gz
chmod +x poc-va-cli
./poc-va-cli