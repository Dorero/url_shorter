# What is it
It is a simple not persistent url shorter service written in go, usage for temporary storage redis.

# How it works
Send your url, get shorted url, click by shorted url, redirect to original url.
\
```POST /url``` - send your url in form data.
\
```GET /url/{id}``` - redirect to original url.

# How setup
```docker-compose up```

# Controversial issues
Storage for url was replaced from postgres to redis, this solution has own advantage and disadvantage, 
and it depends on requirements. I choose redis, because url has a lifetime and no need to remove their from database.