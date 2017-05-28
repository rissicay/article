Technical Test
Article API

The purpose of this exercise is for us to get a sense of how you would approach designing and implementing a simple service, before we get you in for an interview. We’ve tried to avoid tricky algorithmic tests in favor of something that shows how you would organise a codebase.

There is no time limit, but we expect most applicants to spend roughly 2-3 hours working on the test. Once complete please share your repository, forward a zip file of the source code and dependencies, or use a service like Dropbox to share the file.

Feel free to use any language/toolset you like, however, submission written in Go will be looked on favourably. The only requirement is that you can describe how to set it up on a mac so we can see it running.

If you make any assumptions about requirements, or use any online resources to solve a problem, please make note of these somewhere obvious inside the solution (e.g. code comments).

Your solution will be evaluated internally by one or more of your potential co workers. You should expect a response from us within 2 business days.
Requirements

You will be required to create a simple API with three endpoints.

The first endpoint, POST /articles should handle the receipt of some article data in json format, and store it within the service.

The second endpoint GET /articles/{id} should return the JSON representation of the article.

The final endpoint, GET /tags/{tagName}/{date} will return the list of articles that have that tag name on the given date and some summary data about that tag for that day.

An article has the following attributes id, title, date, body, and list of tags. for example:

{
  "id": "1",
  "title": "latest science shows that potato chips are better for you than sugar",
  "date" : "2016-09-22",
  "body" : "some text, potentially containing simple markup about how potato chips are great",
  "tags" : ["health", "fitness", "science"]
}

The GET /tag/{tagName}/{date} endpoint should produce the following JSON. Note that the actual url would look like /tags/health/20160922.

{
  "tag" : "health",
  "count" : 17,
    "articles" :
      [
        "1",
        "7"
      ],
    "related_tags" :
      [
        "science",
        "fitness"
      ]
}

The related_tags field contains a list of tags that are on the articles that the current tag is on for the same day. It should not contain duplicates.

The count field shows the number of tags for the tag for that day.

The articles field contains a list of ids for the last 10 articles entered for that day.
