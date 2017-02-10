---
layout: default
title:  "Quickstart"
---

<h2 id="quickstart">Quickstart</h2> 

Hoverfly comes with a command line tool called **hoverctl**. 

{% highlight shell %}
hoverctl --version
hoverfly --version
{% endhighlight %}

Both of these commands should return a version number. Now you can run an instance of Hoverfly.

{% highlight shell %}
hoverctl start
{% endhighlight %}

Set Hoverfly to **capture** mode to capture an HTTP request and response.

{% highlight shell %}
hoverctl mode capture
curl --proxy http://localhost:8500 http://time.jsontest.com
hoverctl logs
{% endhighlight %}

You have just created a **simulation**. Take a look.

{% highlight shell %}
hoverctl export simulation.json
cat simulation.json
{% endhighlight %}

Set Hoverfly to **simulate mode** and make the same request.

{% highlight shell %}
hoverctl mode simulate
curl --proxy http://localhost:8500 http://time.jsontest.com
hoverctl logs
{% endhighlight %}

The response was returned by Hoverfly. 