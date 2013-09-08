# github bindings

*   create rest calls with request-like library
*   create json parser (`sed`-like)

# github api v3

## authenticate

    $ curl -u <username:password> https://api.github.com/users/defunkt

See basic information on OAuth in the [getting started document][1]

## create repository

    $ curl -u <your_username> -d '{ "name": "<repo>" }' \
      https://api.github.com/user/repos

The resulting repository will be found at `https://github.com/<username>/blog`.
To create a repository under an organization for which you’re an owner, just
change the API method from `/user/repos` to `/orgs/<org>/repos`.

## issues

    $ curl -u <username:password> https://api.github.com/issues

[1]: http://developer.github.com/guides/getting-started/
