# Smart contracts for doge 

- Lead devs: 
    Nicolas Cloutier (@helladdict69 (twitter) Github: godcloutier), 
    
    Finn (Github: fmhc)
    
- Security expert auditor: @dogeconomy
- There will be a section dedicated to all the devs who will participate in this project with their wallet address.

- Language: Backend (C++), Frontend / Smart Doge Wallet (Typescript / GoLang) 


# State of development (June 17th 2021)

We are currently working on a Proof of Concept which will include:

- The ability to store emojis on chain for further parsing 
- The Smart Doge protocol (initial one see below for specs)
- A protocol parser
- RPC Instructions (as a proposal to put on the core dogecoin chain) to make it easier in the future 

# Smart Doge Concept

## What are Smart Contracts on Doge

## Smart Contracts Definition (general)

(Definition coming from Harvard Law school)

“Smart contracts” is a term used to describe computer code that automatically executes all or parts of an agreement and is stored on a blockchain-based platform. As discussed further below, the code can either be the sole manifestation of the agreement between the parties or might complement a traditional text-based contract and execute certain provisions, such as transferring funds from Party A to Party B. The code itself is replicated across multiple nodes of a blockchain and, therefore, benefits from the security, permanence and immutability that a blockchain offers. That replication also means that as each new block is added to the blockchain, the code is, in effect, executed. If the parties have indicated, by initiating a transaction, that certain parameters have been met, the code will execute the step triggered by those parameters. If no such transaction has been initiated, the code will not take any steps. Most smart contracts are written in one of the programming languages directly suited for such computer programs, such as Solidity.

source: [Harvard Law School Forum on Corporate Governance](https://corpgov.law.harvard.edu/2018/05/26/an-introduction-to-smart-contracts-and-their-potential-and-inherent-limitations/ "Smart Contract Definition")

### Why not use an existing Smart Contract Platform like Ethereum (EVM)

Ethereum did pioneer smart contracts on the blockchain and did develop a very useful and amazing language: solidity. Without them a lot of the blockchain space would not be the same today. 

However, Ethereum is more like an all-purpose solution, is managed by a big Fundation (decentralized) and is having features we will not need to be integrated into doge. A bridge is currently talked about between ethereum and doge within the dogecoin developer community. Not much details at this point as to what the bridge would do technically or practically. 

Here are some of the reasons why we decided to develop Smart Doge as an integrated protocol for dogecoin instead of using an existing one: 

- Dogecoin / Doge is a brand, we do not want tokenomics to pollute its space and open the door to "scams" that could dammage the brand like we did see a lot on Ethereum and BSC (Binance smart chain).
- We want to be able to include new, fun and innovative ways to do smart contract "The Doge Way".
- We want to be able to speed up development without going trough lengthy debates and political issues linked to already established project, that could slow down innovation, adoption and reduce doge competitive advantage.


### The Smart Doge Protocol 

This is a protocol proposal and should not be taken as the final version of it. It is a work in progress and we will add to it as it goes:


| 1 BYTE (reg chksum) | 1 BYTE (protocol version | 4 BYTE (optional vm routing) | ~ 506 BYTE (Contract data) |

REG CHKSUM : 00 = regular tx, 01 = segwit, 0a = Smart Doge
PROTOCOL VERSIONS: 

0x0a = NO ROUTING
0x0b = ROUTE WITH INFORMATION FOUND IN THE NEXT 4 BYTES
0x0c = GO TO EMOJIS VM (parse and execute contract code)




# ABOUT GAS FEES (OR THE ABSENCE OF IT)

NONE (AS IN NO GAS FEE THE CONCEPT IS NON EXISTENT)


# ETHICS

Any major change to the core dogecoin chain. If. there is any to be made, it will need to be discuss with the core developer first and be very very minor changes (the impact will need to be determined in the aforementioned discussion first).

Code need to be always safe and tested for security flaws.







# CONTRIBUTIONS

If you would like to help the developers on this project by tipping them collectively you can send doges to the following wallet:
### DTdbhgEE1drrigu7BTWrdh5dJErhod5bFc

All doges received to this wallet will be distributed equally between the contributors of this project and will serve as a motivational purpose, and allow them to put more time on the projects as they can free themselves of their usual daily duties.

For more information, contact HellAddict69 on or Katzenmalen twitter (@helladdict69)



