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

Each component is represented as either a JavaScript class or JavaScript function. Since the release of React 16.8, often referred to as the Hook release. It's been recommended to use functions to represent components over classes, even though classes still work. Going forward it's expected that functions will have more capabilities in React, as well as being better optimized for performance.

```js
'use client';

import React from 'react';
import App from '../src/App';

export default function Home {
    return (
        <App />
    )
}
```

If we look at the return statement it's not JavaScript, it's JSX syntax. The reason it works is that this code project is using **Babel** to transpile this particular file to an **EcmaScript** compliant JavaScript file, and that is what is actually running.

# How React Components Link To Each Other?
Through JSX syntax we formed our component tree, or at least you've seen how from our App components, it's two children components AddForm & TodoItems are liked to it.

```js
export default function App() {
    const [toDoList, setToDoList] = useState([
        {text: "Buy Sugar", id: 1656632538430},
        {text: "Eat Carrots", id: 1656632478827}
    ]);
    .....
    .....
    return(
        <div className="container">
            <div>
                <AddForm addTask={addTask}></AddForm>
            </div>
            <div>
                <TodoItems entries={todoList} deleteItem={deleteItem}>
            </div>
        </div>
    )
}
```

Notice that both AddForm & TodoItems in JSX syntax have attributes with values assigned to them. For example TodoItems passes an entires attribute that is our list of to-dos or todo item entries.
todoList is a javascript array which is initialized by **useState** function call.

**useState** is a react hook, it gives a special way to track data changes over the lifetime of a component, that is as event happens in our component tree like adding a new todo item the code in those event can change what we refer to as React State and that changing data is what makes our components reactive, meaning that they display can change overtime.

Parent components can pass data to their children. The TodoItems referenced in this App component is a perfect example of a parent component, this one named App, passing data to its child component TodoItems. Navigating to our Javascript file, TodoItems.js

```js
import React from 'react';

export default function TodoItems({entries, deleteItem}) {
    return (
        <ul className="mt-3">
            <b className="ms-3">Items:</b>
            {entires.map(({ id, text }) => {
                return (
                    <li onClick={() => deleteItem(id);}>{{text}}</li>
                )
            })}
        </ul>
    );
}
```

Notice this components have been passed entries & deleteItem values from App component. Instead of using props we are using destructuring props, which is more clear and readable.

We can write Javascript inside JSX (like entries loop here) and then that javascript can return more jsx code (like li items here).

# Oneway Databinding
React was created based on the frustrations with the web app frameworks that implement two-way data binding with complex MVC architectures that use HTML templates as the precompiled source for data views.

The way other frameworks with two way data binding works is that:
They typically compile HTML templates into a View and then an associated data model needs to be synchronized with that view, that is when the data model changes the view in the browser needs to be updated. That creates a tight loop between the model and the view that is continually in need of updating as either the view changes or the model changes.

The design of React is different. There is not a tight coupling between the Model & the view which leads to React being one-way data binding, instead of two. Rather than using HTML templates to define HTML views, React instead, uses components, where the components not only have the HTML view embedded in them but they also have the logic for how to update those views. React components essentially all linked together to create a single virtual DOM, which is the core of a React App.

That Virtual DOM is completely self contained in JavaScript, meaning that any changes to the data model in a React APP only update the virtual DOM and don't have to interact with the App's Browser DOM, JavaScript to Browser DOM updates are very slow relative to JavaScript updates to its own objects that is the Virtual DOM in this case. Once all the Model changes are complete, React compares a prior copy of the Virtual DOM from before updates were processed with the new Virtual DOM and then figures out through a merge process called **reconciliation** what DOM elements need to be updated.