# Gator

Gator is a cli RSS Aggrigator. You can register and log in different users. Each of them can add different feeds to fetch/follow. A user can follow multiple feeds and get posts from them.

## Requirements

- Postgres v15+
- Go v1.25+

## Installation

Use the go install command to install the gator cli.

```bash
go install
```

## Usage

You will need to create a config file in your home folder, `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://<username>:<password>@localhost:5432/gator",
  "current_user_name": ""
}
```

Then you can use the following commands:

- `login <user_name>` - login as a user
- `register <user_name>` - register a new user
- `users` - list all users
- `agg` - start aggregating feeds
- `addfeed <feed_name> <feed_url>` - add a new feed to the program
- `feeds` - list all available feeds
- `follow <feed_url>` - follow a feed
- `unfollow <feed_url>` - unfollow a feed
- `following` - list all the feeds current user follows
- `browse` - browse latest posts

### Example usage:

```bash
$ gator login aditya
User switched successfully!

$ gator following
Feeds followed by user aditya:
 * TechCrunch
 * Hacker News
```
