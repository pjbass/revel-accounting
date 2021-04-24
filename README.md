# Accounting

A simple accounting app built using Revel and GORM.
Allows adding and deleting assets and liabilities, and calculates total assets,
total liabilities, and net worth.

## Start the web server:

   revel run myapp

## Layout

The application consists of a single page where assets and liabilities can be
added and deleted. There are two additional endpoints: one for adding assets/liabilities
and one for deleting them. All of these actions are provided by the App Controller.
