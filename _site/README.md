# Hoverfly.io site

This is a static site generated with Jekyll.

For local development:

```
docker run --rm -i -t \
           --name=jekyll-builder \
           --volume=/$(pwd):/srv/jekyll/ \
            -p 4000:4000 \
            jekyll/jekyll:builder \
            jekyll serve -w
```            

For each Hoverfly release, the `hf_version` property in `_config.yml` must be updated.
