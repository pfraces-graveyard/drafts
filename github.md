# github bindings

*   create rest calls with request-like library
*   create json parser (`sed`-like)
*   read [hub][5] syntax

# github api v3

## authenticate

    $ curl -u <username:password> https://api.github.com/users/<username>

See basic information on OAuth in the [getting started document][1]

## create repository

    $ curl -u <username:password> -d '{ "name": "<repo>" }' \
      https://api.github.com/user/repos

The resulting repository will be found at `https://github.com/<username>/blog`.
To create a repository under an organization for which you’re an owner, just
change the API method from `/user/repos` to `/orgs/<org>/repos`.

## get assigned open issues

    $ curl -u <username:password> https://api.github.com/issues

## get created open issues

    $ curl -u <username:password> https://api.github.com/issues?filter=created

## create issue

    $ curl -u <username:password> -d '{
      "title": "<title>",
      "body": "<issue>",
      "labels": ["<labelname>"] }' \
      https://api.github.com/repos/<username>/<repo>/issues

# further documentation

*   [manage labels][2]
*   [delete repo][3]
*   [more on issues][4]


[1]: http://developer.github.com/guides/getting-started/
[2]: http://developer.github.com/v3/issues/labels/
[3]: http://developer.github.com/v3/repos/#delete-a-repository
[4]: http://developer.github.com/v3/issues/
[5]: http://hub.github.com/
