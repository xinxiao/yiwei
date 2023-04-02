# Yiwei
-----------

### Intro
Yiwei is the sound of the Chinese word 一维, which means one-dimension. The goal of this 
database is to allow people to annotate data points in the 1-dimensional time series, so 
that they  could drill down by those labels and detect anomalies in specific dimensions 
a bit easier; plus aggregation is also easier in a way :).

### Incentive
I have been working on reliability throughout my short career, with this immersion I developed
this almost-reflex intention to evaluate and comapre reliability solutions, such as monitoring 
and alerting, so that I could help myself better understand how I may construct a reliable 
software system quickly and efficiently.

When I was with this famous web search engine, they offered a very powerful dashboard solution 
that allowed oncalls to easily drill down monitoring series by dimensions such as regions deployed, 
versions of binary, success rate of calls to different dependencies... The tool was so powerful 
that with it I could always narrow things down after a few clicks. However, after joing this 
company that inspired the social network movie, I realized that they adopted a distinct monitoring
paradigm where that FWSE's split up feature is no longer accessible. Therefore, I started to
wonder if it might make sense to whip up a quick solution that:
1. Still works decently as a time series store;
2. Allows people to somehow annotate the datapoints with some metadata so when they need to drill 
   down, they can.

So ladies and gentlemen, let's welcome Yiwei, a database that was implemented with all the assumptions
a time-series store could make, meanwhile added a hint with annotations so that people can cut the 
series to a specific dimension if they'd like to do so.

### Design
Like all time series, data points in Yiwei has the 2 primary attributes:
1. A timestamp used as the index/primary key;
2. A long float number to represent the desired reading.

Except these two, people can also choose to annotate them with a series of `label` while pumping 
the data points to Yiwei. Each label is essentially a pair of strings, `key` and `value`. While 
querying people can make a filter based on these `label` to get acquired the desire subsets of 
data within the slected time range. Currently there are 2 types of query interfaces suported,
batch and stream.

Yiwei maintains a very simple set of storage engine with the assumption that:
1. Timestamp index is always increasing;
2. Accepted data are immutable.

With this assumption, the storage layer is designed similar as a 2-layer append-only log. The 
fundamental layer is called [`Page`](/database/page/page.go), where all data points are sorted
consecutively, and sorted by index, since time doesn't go back (sometimes I wish). Each page
is assigned with a UUID, and once it's full it will request the layer on top to create a new 
page and remeber the ID of the immediate successor.

The layer on top of `Page` is named [`Series`](/database/series/series.go). `Series` organizes
and indexes `Pages`; it maintains a consecutive list of `Pages` annotated with the first index
of themselves, so based on this we could easily find out the range of data within that `Page`.

Since indices in both layers are on sorted, it's relatively straight-forward to  query: first 
do a binary search in series index to locate the page with the largest first index that's 
smaller than the specified start, and in that page again do a binary search to locate the entry 
with the smallest index that's >= the specified start. Then from that start, follow the page chain 
until we find the last value that's <= the specified end.

#### Action Items (who knows when I can get to this :( )
- [ ] Implement series operation (I know, but I started yesterday night so what do you expect).
- [ ] Support index on `label` key.
- [ ] Speed up persistene IO by caching some files.
- [ ] **HIGHLY UNLIKELY**: Write unit tests.
- [ ] Come up with more.

### Dev Guide
*I really don't expect anyone to be reading this but here we go :p.*

This project is built with [Bazel](https://bazel.build/), so please follow the guide on
Bazel's site to complete the necessary steps to have the tool installed on your machine.

The binary of Yiwei service locates at the folder [`yiwei`](/yiwei), and you may build the binary by
running:
```bash
$ bazel build //yiwei
```

Yiwei currently situates its interface on top of [gRPC](https://grpc.io/), going over the basics
of this RPC framework may help you better understand how Yiwei connects its pieces all together. 
The manifest that defines Yiwei's interface is [here](/proto/database.proto).

It's also easy to spawn up an instance of Yiwei; simply run:
```bash
$ bazel run //yiwei
```

However, if you may notice that the instance quickly fail due to permission. By default Yiwei
writes all of its data files to `/var/lib/yiwei`, so for an unpreviledged user please consider
overriding the data directory to an accessible place by:
```bash
$ bazel run //yiwei -- --data_dir=<desire_location>
```

Yiwei uses many [command line flags](https://gobyexample.com/command-line-flags) to pass
parameters to runtime. To understand what parameters are involved, please run:
```bash
$ bazel run //yiwei -- --help
```

### Closing
That will be all for now my friends; if I could get some time between shoveling ads 
to your Facebook/Instagram feeds I'll come back with more ideas/updates :)

P.S.: To remeber a beatiful one

-- @xinxiao, 04/01/23
