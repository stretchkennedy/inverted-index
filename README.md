# inverted-index
Inverted index for full-text search written in go and queriable via gRPC. 

This probably goes without saying, but _do not actually use this in production_.
It's a toy project created in a weekend and it's missing very important 
features, like thread safety for writing to the map, persistence, configurability, 
and probably a lot more. 
