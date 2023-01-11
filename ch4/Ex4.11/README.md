# Issue CRUL (Create, Read, Update & Lock) CLI made with Go!  
This is a simple CRUL tool for creating, reading, updating and locking issues using the [github API](https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#about-issues).

## Setup
First you need to [generate a github personal access token](https://docs.github.com/en/enterprise-server@3.4/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) to be able to write to your own issues on your repos.  
After you get the token, create a `.env` file on the root of this project and add the following line.  
```
GITHUB_TOKEN=<YOUR-ACCESS-TOKEN>
```
do not include "<>" on your token.  

## How to use it?
Use the go command line commands:  
```
$ go build  
or  
$ go run .
```
if you dont give any parameter the tool tells you how to use it:  
```
$ go run .
2023/01/10 10:05:07 Usage:
 -./bin <CRUL>
--------------------------------
 - <CRUD>: e.g. "create"|"read"|"update"|"lock"  
exit status 1
```

### Create an issue  
Just follow the steps, for example:
```
$ go run . create
Write the name of the repo: TomoDiagFrontend
Write the name of the repo's owner: webtaken
Write the title of the issue: my issue
What's the issue?: body of my issue
Endpoint https://api.github.com/repos/webtaken/TomoDiagFrontend/issues
Issue created successfully 😄
{
    "Number": 5,
    "html_url": "https://github.com/webtaken/TomoDiagFrontend/issues/5",
    "title": "my issue\n",
    "state": "open",
    "User": {
        "Login": "webtaken",
        "html_url": "https://github.com/webtaken"
    },
    "created_at": "2023-01-10T14:45:36Z",
    "body": "body of my issue\n"
}
```
### Read an issue  
Just follow the steps, for example:
```
$ go run . read
Write the name of the repo: TomoDiagFrontend
Write the name of the repo's owner: webtaken
Write the issue number: 5
Endpoint: https://api.github.com/repos/webtaken/TomoDiagFrontend/issues/5
Issue readed successfully 😄
{
    "Number": 5,
    "html_url": "https://github.com/webtaken/TomoDiagFrontend/issues/5",
    "title": "my issue\n",
    "state": "open",
    "User": {
        "Login": "webtaken",
        "html_url": "https://github.com/webtaken"
    },
    "created_at": "2023-01-10T14:45:36Z",
    "body": "body of my issue\n"
}
```
### Update an issue  
Just follow the steps, for example:
```
$ go run . update
Write the name of the repo: TomoDiagFrontend
Write the name of the repo's owner: webtaken
Write the issue number: 5
Write the new title of the issue: My new title
Write the new body of the issue: My new issue body
Endpoint: https://api.github.com/repos/webtaken/TomoDiagFrontend/issues/5
Issue updated successfully 😄
{
    "Number": 5,
    "html_url": "https://github.com/webtaken/TomoDiagFrontend/issues/5",
    "title": "My new title\n",
    "state": "open",
    "User": {
        "Login": "webtaken",
        "html_url": "https://github.com/webtaken"
    },
    "created_at": "2023-01-10T14:45:36Z",
    "body": "My new issue body\n"
}
```
### Lock an issue  
Just follow the steps, for example:
```
$ go run . lock
Write the name of the repo: TomoDiagFrontend
Write the name of the repo's owner: webtaken
Write the issue number: 5
Select one of the following options for locking this issue:
1) "off-topic"
2) "too heated"
3) "resolved"
4) "spam"
3
Endpoint: https://api.github.com/repos/webtaken/TomoDiagFrontend/issues/5/lock
Issue Locked successfully 😄
```
## Thanks & Resources 😄  
Thanks to me (XD) and internet for starting this journey with this great programming language. I coded this little challenge from the book [The Go Programming Language](https://www.gopl.io/), hope you liked the little project.