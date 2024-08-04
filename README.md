# Interacting with the blockchain

- Create a Resource:  
`myblockchaind tx myblockchain create-resource title body --from alice --chain-id myblockchain`

- View a Resource:  
`myblockchaind q myblockchain show-resource 0`

- List All Resources:  
`myblockchaind q myblockchain list-resource`

- Update a Resource:  
`myblockchaind tx myblockchain update-resource "Updated title" "Updated body" 0 --from alice --chain-id myblockchain`

- Delete a Resource:  
`myblockchaind tx myblockchain delete-resource 0 --from alice  --chain-id myblockchain`

## Consensus-Breaking Change

### What is Breaking Consensus?

Breaking consensus means making a change to the blockchain that is not backward-compatible. Nodes running the old version of the software will not be able to reach consensus with nodes running the new version, leading to a fork.

### Change That Breaks Consensus

`id` field in `resource.proto` file changed from `uint64` to `string`, and packages are updated to reflect this change.

### Why This Change Breaks Consensus

The change from `int` to `string` for the `id` field in the `Resource` message changes the way data is stored and interpreted. Nodes running the previous version will interpret the `id` as an integer, while the new version interprets it as a string. This discrepancy will cause consensus failure between different versions.
