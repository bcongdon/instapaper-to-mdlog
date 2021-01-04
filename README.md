# instapaper-to-mdlog

Small CLI tool to keep a running log of archived Instapaper articles.

I built this to capture my Instapaper reading history in [Obsidian](https://obsidian.md/), but this tool outputs generic markdown, so can work with a variety of tools.

There's also nothing Instapaper specific about this, so it could also be used for other services that offer RSS feeds to indicate activity.

## Installation

```sh
$ go install github.com/bcongdon/instapaper-to-mdlog
```

## Format

`instapaper-to-mdlog` produces and updates a "dated log" in Markdown list format of archived Instapaper items. The look looks like the following:

```md
# January 3, 2021

- [Some Article](http://example.org/article1)
- [Another Article](http://example.org/article2)

# January 2, 2021

- [Yet Another Article](http://example.org/article3)

# January 1, 2021

- [Yet Yet Another Article](http://example.org/article4)
```

Note that the date of each article is _not_ the publish/update date of the article, but rather the date that the article was discovered by `instapaper-to-mdlog`. This is because `instapaper-to-mdlog` is designed to capture the day you read the article (e.g. archived it) instead of the day that the article was created.

## Usage

First, go to your [Instapaper Archive](https://www.instapaper.com/archive), click on your username/email in the upper-right corner, and then click "Download > RSS Feed". Copy the URL of the resulting feed, and use it as the input to `-feedURL`.

You can set where the log is updated with `-out`.

```
$ instapaper-to-mdlog -feedURL <FEED_URL> -out out.md
```

`instapaper-to-mdlog` is designed to be used with `cron`. For example, by updating your reading log every hour:

```sh
0 * * * * /path/to/instapaper-to-mdlog -feedURL $FEED_URL -out /path/to/out.md
```
