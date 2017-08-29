# The goal
Need to download GP all time archive. It looks like all the articles can be accessed using url in form http://www.gp.se/1.%07d.
So to get all the pages we can iterate through 0 to 9999999.
As a second solution we can use BFS, but some articles can be not accessible.

# Data Source
Meaningful data in the source html

* <div class="article__body__richtext container ">. This is the main body

<div class="article__body__richtext container ">
      <p>Det var i samband med tr&auml;df&auml;llnings som olyckan intr&auml;ffade i N&auml;rebo strax utanf&ouml;r LIdk&ouml;ping.</p> 
        <p>– Tr&auml;det sprack och han tr&auml;ffades och f&ouml;ll ner i en b&auml;ck, s&auml;ger Mikael Bengtsson som &auml;r insatsledare p&aring; R&auml;ddningstj&auml;nsten i V&auml;stra Skaraborg. </p> 
        <p>Mannen var medvetande under hela r&auml;ddningsinsatsen men har f&aring;tt skador p&aring; flera st&auml;llen p&aring; kroppen enligt r&auml;ddningstj&auml;nsten.</p>
</div>

* datetime. We shall use only day.
      <time datetime="2016-02-26">
        13:51 - 26 feb, 2016
      </time>

* additional info. We can use category
<div class="article__head ">
  <span id="article-data-1.532" style="display: none;" category-main="nyheter" category-sub="vastsverige" premium="false" access="true" native="false" campaign-id="" source="Newspilot" top-element-type="image" top-element-id="1.531"></span>

* article title
  <!DOCTYPE html>
  <html lang="sv">
    <head>
    <title>Föll i bäck efter att ha träffats av träd | Göteborgs-Posten - Västsverige</title>


### The output should be a list of files.

* Name format:
gpdumpId.txt

* Contents format
link
yyyy_dd_mm
title
category
body

### TODO: 
+ have all articles in one file
+ swedish encoding?