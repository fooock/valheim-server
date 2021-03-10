
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
	ansible-galaxy collection install -f -r valheim-ansible/requirements.yml
	ansible-galaxy role install -f -r valheim-ansible/requirements.yml

ansible-install: ansible-dependencies
ifndef TARGET_IP
$(error TARGET_IP is not set. Please provide a valid IP/host)
endif
	ansible-playbook -i $(TARGET_IP), -u $(TARGET_USER) valheim-ansible/playbook.yml
