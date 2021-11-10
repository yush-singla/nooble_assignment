# nooble_assignment

To run the server use
```
go run server.go
```

The server will spin up on port 8080
with a graph ql api where the mutations/queries can be used 

The following queries/mutations can be used (by changing the ids ofc for individual documents)
Uncomment the following to use them, and then play

```
# {
#   creators{
#     name,
#     id
#   }
# }

# {
#   audios{
# 		id,
#     title,
#     description,
#     creatorId,
#     category,
#     creator{
#       name,
#       email,
#       id
#     }
#   }
#   creators{
#     name,email
#   }
# }

# mutation{
#   createAudio(input:{title:"1testing1",description:"1testing",category:"1boring",CreatorId:"6129484611666145821"}){
#     description,
#     creator{
#       name,email
#     },
#     title
#   }
#   # createCreator(name:"abc1",email:"abc1@gmail.com"){
#   #   name,
#   #   email,
#   #   id
#   # }
# }

# {
#   # creator(input:"5577006791947779410"){
#   #   id,
#   #   name,
#   #   email
#   # }
#   audio(input:"5577006791947779410"){
#     title,
#     description,
#     id,
#     creator{
#       name,
#       email
#     }
#   }
# }
```

The api provides functionality for 
1. Creating a new creator
2. Creating a new audio 
3. Get list of all creators/audios
4. Get list of creator/audio based on audio

For each audio short we can fetch, the creator of the audio short as well 
