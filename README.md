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

