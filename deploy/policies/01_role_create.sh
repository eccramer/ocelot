#!/usr/bin/env bash
echo "writing ocelot policy"
vault policy write ocelot ocelot.hcl
echo "writing werker policy"
vault policy write werker-deploy-restricted werker_deployer.hcl
