# Topictool

***CLI Tool to manage topic labels on Github repositories***

![](https://img.shields.io/badge/Go-1.17%2B-blue)
![](https://img.shields.io/github/issues/cbrgm/topictool)
![](https://img.shields.io/github/license/cbrgm/topictool)

## Installation

```
go get github.com/cbrgm/topictool
```

or

```
git clone git@github.com:cbrgm/topictool.git && cd topictool
go mod vendor && make
```



## Usage

You'll need a ***Github Personal Access Token*** to use this tool. Create one for your user (Settings -> Developer Settings -> Personal Access Token) and grant ***read/write access for repositories*** to it. In case you want to modify private repositories of an ***organization*** please authorize SSO. 

```bash
export GH_ACCESS_TOKEN=<your-token>
```

then run the `topictool`.

```
Usage: topictool <subcommand> <search pattern> <topic labels...>

Replace, add or remove topic labels from multiple Github repositories

Subcommands:
    - replace   - replaces all existing topic labels with new ones
    - add       - adds topic labels to existing ones
    - rm        - removes topic labels from existing ones
    
Search Pattern:
    Searches repositories via various criteria.
    See Github docs: https://docs.github.com/en/free-pro-team@latest/rest/reference/search/#search-repositories

Topic Labels:
    A list of strings representing topic labels

```

### Examples

***Add a labels `foo` and `bar` to the repository `cbrgm/topictool`***

```bash
topictool add "cbrgm/topictool" "foo" "bar"
```

```bash
Repository Name	Topics	
 ---		---	
 cbrgm/topictool	

 Add labels [foo,bar] to 1 repositories? [y/n/q]:
 
 > y
 
 Done!
```

***Add a labels `foo` and `bar` to all repositories of user `cbrgm`***

```bash
topictool add "user:cbrgm" "foo" "bar"
```

```bash
Repository Name			Topics								
 ---				---								
 cbrgm/telegram-robot-rss		bot,messenger,python,rss,rss-feed,rss-reader,telegram-bot			
 cbrgm/go-t			        cli,command-line,go,golang,twitter,twitter-api,twitter-client		
 cbrgm/clickbaiter			advertising,clickbait,generation,generator,go,golang,useless		
 cbrgm/terraform-k8s-hetzner		bash,hetzner,hetzner-cloud,kubernetes,terraform				
 cbrgm/authproxy			api,authentication,client,kubernetes,middleware,token,webhook		
 cbrgm/kubernetes-rbac-groups		authorization,cluster,kubernetes,rbac	
 
Add labels [foo,bar] to 6 repositories? [y/n/q]:

> y

Done!
```

***Replace all labels with `foo` and `bar` for all repositories of user `cbrgm` already having a topic `foo`***

```bash
topictool add "user:cbrgm topic:foo" "foo" "bar"
```

```bash
Repository Name			Topics								
 ---				---								
 cbrgm/telegram-robot-rss		bot,messenger,python,rss,rss-feed,rss-reader,telegram-bot,foo			
 cbrgm/go-t			        cli,command-line,go,golang,twitter,twitter-api,twitter-client,foo		
 cbrgm/clickbaiter			advertising,clickbait,generation,generator,go,golang,useless,foo		
 cbrgm/terraform-k8s-hetzner		bash,hetzner,hetzner-cloud,kubernetes,terraform,foo				
 cbrgm/authproxy			api,authentication,client,kubernetes,middleware,token,webhook,foo		
 cbrgm/kubernetes-rbac-groups		authorization,cluster,kubernetes,rbac,foo	
 
Replace all existing topic labels with [foo,bar] in 6 repositories? [y/n/q]:

> y

Done!
```

## Contributing & License

Feel free to submit changes! See the [Contributing Guide](https://github.com/cbrgm/contributing/blob/master/CONTRIBUTING.md). This project is open-source and is developed under the terms of the [Apache 2.0 License](https://github.com/cbrgm/topictool/blob/master/LICENSE).
