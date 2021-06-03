# Smart contracts for doge 

- Lead dev: @helladdict69 (twitter) Github: godcloutier 
- Main dev: @dogeconomy (twitter)

- Language: GoLang 1.16.4 (latest), Rust (networking layer)

## Goal: 

Make smart contracts available as "plugins" or "modules" on the doge blockchain. Conditions #1 to fulfill the goal: Add a data field to dogecoin core code (8 bytes) to allow "verified" apps on the chain to run as plugins.

If condition #1 is not met: By Proxy

Version 0.0.1 will be by proxy by default

An issue / PR will need to be made to confirm condition #1 (if accepted by the dev community)

# IMPORTANT

We will not include :

- Tokenization (no use case for it with a plugin system)
- Value manipulation (setting a value to an asset other than based on dogecoin actual value)

## ABOUT GAZ FEES (OR THE ABSENCE OF IT)

Gaz fees (if there is to be fee eventually it will be decided by the miners, however it will greatly depends on how we chose to integrate the project in the decentralize environment)


## ETHICS

Any major change to the core dogecoin chain. If. there is any to be made, it will need to be discuss with the core developer first and be very very minor changes (the impact will need to be determined in the aforementioned discussion first).

Code need to be always safe and tested for security flaws.

## TO BE DEVELOPED IN THIS VERSION

- Communication Proxy with the chain (a la rosetta api? protocol buffers? something better?)
- DOGITY : The language to code the contracts
- An ipfs connector to decentralize assets (contract descriptions, file types, etc) @Dogeconomy idea
- A secure way to anonymize transactions at high level (leaving low level code untouch) with key exchange to secure possible sensitive content (use case: medical field per example)

..... ADD MORE FEATURES HERE .....
