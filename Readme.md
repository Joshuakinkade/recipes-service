# Recipes Service

The Recipes Service is an API for managing a collection of recipes.

- experimenting with architecture and project layout
- written using Copilot so I can focus on the bigger picture parts of the project.

## Data Model

I created this handwritten diagram over a year ago. The purple outlines define boundaries of different services. The yellow outline defines features I brainstormed and modeled, but decided not to implement right away. I also have some notes about how I'm going to run database queries to make sure my model won't force me into any bad queries. This thinking led me to put the recipe's main photo in the recipe table as well as the photos table to avoid doing any joins while listing a user's recipes.

![Data Model Diagram](./Data%20Model.png)

## Unit Conversion

The most programatically interesting part of this project is the unit conversion. I want to allow the users to change the units that the recipes are saved with and change the units that are displayed while viewing a recipe. Also want to show the units in a clean and concise format, so I need to figure out how to keep long ugly decimals from showing up.
