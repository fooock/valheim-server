
#                  _ _     _
#                 (_) |   | |
#   __ _ _ __  ___ _| |__ | | ___
#  / _` | '_ \/ __| | '_ \| |/ _ \
# | (_| | | | \__ \ | |_) | |  __/
#  \__,_|_| |_|___/_|_.__/|_|\___|
#
TARGET_IP :=
TARGET_USER := root

ansible-check:
	ansible-playbook valheim-ansible/playbook.yml -i valheim-ansible/game-role/tests/inventory --syntax-check

ansible-dependencies:
	ansible-galaxy collection install -r valheim-ansible/requirements.yml
	ansible-galaxy role install -r valheim-ansible/requirements.yml

ansible-install: ansible-dependencies
ifeq ($(strip $(TARGET_IP)),)
	$(shell echo "TARGET_IP is not set. Please provide a valid IP/host")
else
	ansible-playbook -i $(TARGET_IP), -u $(TARGET_USER) valheim-ansible/playbook.yml
endif

##                    _
##   _ __   __ _  ___| | _____ _ __
##  | '_ \ / _` |/ __| |/ / _ \ '__|
##  | |_) | (_| | (__|   <  __/ |
##  | .__/ \__,_|\___|_|\_\___|_|
##  |_|

CLOUD := 
PACKER_BIN := /usr/local/bin/packer

packer-validate:
	$(PACKER_BIN) validate valheim-packer/hcloud.json
	$(PACKER_BIN) validate valheim-packer/vagrant.json

packer-run:
ifeq ($(CLOUD), hcloud)
	$(PACKER_BIN) build valheim-packer/hcloud.json
else ifeq ($(CLOUD), vagrant)
	$(PACKER_BIN) build valheim-packer/vagrant.json
endif
