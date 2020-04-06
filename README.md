# covid19

Together We Can Make a Difference. Become a contributor today.

Our Websites:

https://outbreak.cc

https://pandemic.ai

Our Twitter
https://twitter.com/pandemic_ai

TODO List:
- Add US Maps and details about cases/deaths in each state
- Add UK Maps and case details
- Update timeline with details
  (Automate this using GO or Python scripts with web scraping)
- Add Maps and Case details for all affected countries (at least top 20 countries)
- Add charts, heat maps, and tickers for all major sectors and stock markets to track the impact of the coronavirus outbreak on the global economy
  (Automate this)
- Develop dashboard Web app based on React framework to track the coronavirus outbreak in each country

Software Projects:

1) Develop a frontend that converts JSON data file into HTML table. 
The frontend should allow user to add rows to the table, move rows up or down, and enter data in the fields.

Tabulator library provides most of the needed funcitonalities.
http://tabulator.info/examples/4.5#editable

2) Dashboard Apps based on React framework
Simialr to what they have in Hongkong and Singapore
https://chp-dashboard.geodata.gov.hk/covid-19/en.html
https://www.moh.gov.sg/covid-19

3) Webscraping statistics for coronavirus cases and automate the process of data collection using Python or GO scripts.

Data projects:
1) Update Timeline with details
2) Update Youtube Playlist per country (This can be automated)
3) Update "must read" section (This can be done through a Webform)

We need
- Frontend developers (Javacript, Typescript with React Framework)
- Backend developers (Python Flask) with data/web scrapping experience
- UI developer
- People who know how to create inforgraphics using Adoble Illustrators and Powerpoint
- People who have the time to update the data in JSON files
- Peolpe who can help with fundraising

You can edit JSON files with this online editor https://json-csv.com/editor
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

If you are new to git, please read this cheat sheet first.

https://github.github.com/training-kit/downloads/github-git-cheat-sheet.pdf
https://rogerdudler.github.io/git-guide/

If you are new to python, please complete this tutorial first

https://docs.python.org/3/tutorial/

If you are new the Javascript or Jquery, please complete those tutorials first

https://javascript.info/

https://learn.jquery.com/

If you are new to Visual Studio Code, please complete this tutorial first

https://flaviocopes.com/vscode/

