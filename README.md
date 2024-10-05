# discnet
Service that queries Discogs API to build a weighted graph of collaborations between musicians

# Build and run

```console
git clone https://github.com/njvanhaute/discnet.git
cd discnet
```

At this point, you'll need to create a .envrc file exporting DISCOGS_API_URL, DISCOGS_API_KEY, DISCOGS_API_SECRET, and DISCNET_USER_AGENT.

Then, just run

```console
make run/api
```

or 

```console
make wrun/api
```

if you want live reload via wgo.
