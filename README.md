# The Brief

You have selected the "Global Ethical Fund" option with your pension provider, believing that your pension contributions would consequently be invested in a range of different companies around the world, all of which are run ethically.

Recently you read a newspaper article that bought into question just how ethical some of the investments in the "Global Ethical Fund" may be, in particular a company called "GoldenGadgets" is rumoured to be using gold mining practices that have a negative impact on the local environment and their staff welfare record is similarly dubious. You decide to investigate for yourself, and set about finding out how your pension contributions are invested, and in particular if your own pension includes a holding in "GoldenGadgets".

You discover that your pension is actually invested in a number of different ***funds***. Each of these funds comprises shares in some well-known large companies as well as some other funds. You dig further and see that this pattern repeats itself indefinitely, where a fund may contain regular shares and/or other funds.

You would like to find out the full list of shares that are held by your pension, either directly or indirectly through one or more funds to see if "GoldenGadgets" is amongst them.

* A fund may be invested in any number of different company shares (e.g. Apple, Microsoft, Diagio, etc)
* A fund may also be invested in any number of other funds (e.g. FTSE 100 Fund, or Japan Small Cap Fund).
* There is no limit to how deeply nested funds may become
* You may assume that there are no "cycles" in the tree of funds, so if Fund A includes Fund B, Fund B will not include Fund A.
* You may also assume that the names of funds and shares are unique within the input file

The task is to read a list of funds from a file, formatted as defined below, and to output the names of all the different company shares held within the "Ethical Global Fund". Do not include the names of any intermediary funds in this list, just the names of the companies that are ultimately held.

For each company you hold, you should use the weight value for each holding to calculate and output, the total proportion of your investments that are held in that company. The `weight` property in the sample tells you what proportion of the total value of the parent fund that the holding represents. Note that some funds may invest in the same companies as other funds, but as the pension holder, I just want to see a consolidated list of the companies in my portfolio, and my total exposure to each one.

Does your pension, represented by the Global Ethical Fund, include GoldenGadgets? If so what proportion of your pension is currently invested in the company?