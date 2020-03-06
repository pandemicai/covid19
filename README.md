# covid19

Our discord server:
https://discordapp.com/invite/FPQvNR

Our Websites:
https://outbreak.cc
https://pandemic.ai

Our Twitter
https://twitter.com/pandemic_ai

We need
- Frontend developers (Javacript, Typescript with React Framework)
- Backend developers (Python Flask) with data/web scrapping experience
- UI developer
- People who know how to create inforgraphics using Adoble Illustrators and Powerpoint
- People who have the time to update the data in JSON files

You can edit the data in Google sheets

Live Update:
https://docs.google.com/spreadsheets/d/1yRGeLf5qUcKxq_zcv8K0zyDImNg6t7nAtgULHc-UyP4/

Summary Report:
https://docs.google.com/spreadsheets/d/1pg4CW6ipf5FwakLkwneqdbZxBsdSpr2UflrjrKGKHgA/

Timeline:
https://docs.google.com/spreadsheets/d/1g_1UqT2RPAnWjWylTco7edUxfQxHZu6gT2-fuCNPvaw/

Please request access.

You can also edit JSON files with this online editor https://json-csv.com/editor
You can convert JSON to CSV at https://json-csv.com. Then you can edit the file in Excel. Once you are done, you can convert the CSV back to JSON at https://json-csv.com/reverse

Try it.  https://raw.githubusercontent.com/pandemicai/covid19/master/json/covid19_live_update.json

Use https://www.iban.com/country-codes to look up the two letter country code if you need to add a country.

If you are a programmer, I recommend using Visual Code Studio as the editor. It is free. You can download a copy from here:
https://code.visualstudio.com/

Important things:
- You can not have doule quotes "" in Json key/value pairs. Replace the double quotes with a single quote.
- You can not break json string into multiple lines.
- Using https://jsonlint.com/ JSON validator to validate the json files.
- Don't introduce new key/value pairs. Don't delete existing keys/value pairs.
- You can update the values in each row of the JSON files, add and delete rows
- Each row needs to end with a comma ",", except for the last row

For live update, make sure the rows are sorted by number of confirmed cases, and update the URL reference for the new reported cases (google search for the link, no website with a paywall. Trusted source only, take a look existing JSON file for commonly used sources)
Use  https://www.worldometers.info/coronavirus/ as your data source.

For timeline update,  use https://bnonews.com/index.php/2020/02/the-latest-coronavirus-cases/ as your data source. You use google search for URL links.

For map update, use this as a template and create maps for your own country.
https://www.google.com/maps/d/viewer?mid=1yCPR-ukAgE55sROnmBUFmtLN6riVLTu3
