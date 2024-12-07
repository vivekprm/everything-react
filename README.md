# What is React
Frontend javascript library created by facebook.

According to builtwith.com there are 7 million websites that use React.

# React Component Model
At the core of the react is component model that you as a Javascript Developer, create and update. As updates are made to the component model in Javascript, the user's browser is instantly updated.

As a developer we typically think of a collection of components as a tree. All react apps have a single component at the top of that tree that we refer to as the root component of our app.

All components can have 0, 1 or more children. What makes React apps both intuitive to program, as well as hugely performant is that inside components in our component tree, we trigger JavaScript events that can update one or more components anywhere in that tree.

Following that React Engine, blazingly fast, updates the browser UI to reflect the new updated component tree.

<img width="1215" alt="Screenshot 2024-12-07 at 12 19 57 PM" src="https://github.com/user-attachments/assets/2a4c4003-de68-4f08-b592-5fbc87ab1ca4">

For example, let's type a new to-do item in our input text field which is a child of AddForm component, then press it's sibling component the AddItem button. What happens in our component tree is that the AddItem button sends whatever value was typed into the input field component upto the AddForm component and in turn sends that value up to the App component. That App component automatically sends the new item to the TodoItems component that then creates a new child Item component. That becomes the third Item in our todo list.

<img width="1212" alt="Screenshot 2024-12-07 at 12 34 14 PM" src="https://github.com/user-attachments/assets/fcfdcaa7-4a63-4b2a-be7a-e2ebd37026a0">

Once third Item is added our browser immediately updates to reflect the new state of our component tree which is that there are three items now in the todo list. 
UI updates practically simultaneously to the tree being updated. Typically when a component tree update happens, the UI reflects that change in less time than it takes for a single frame of the browser to be displayed on the monitor (i.e. 30 or 60 fps). That's the primary reason that React Apps run so well in our browsers.
