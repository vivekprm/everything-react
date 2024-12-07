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

<img width="1162" alt="Screenshot 2024-12-07 at 3 05 54 PM" src="https://github.com/user-attachments/assets/e4f72e69-dcb1-4c3d-9743-d90970bdca35">

That Virtual DOM is completely self contained in JavaScript, meaning that any changes to the data model in a React APP only update the virtual DOM and don't have to interact with the App's Browser DOM, JavaScript to Browser DOM updates are very slow relative to JavaScript updates to its own objects that is the Virtual DOM in this case. Once all the Model changes are complete, React compares a prior copy of the Virtual DOM from before updates were processed with the new Virtual DOM and then figures out through a merge process called **reconciliation** what DOM elements need to be updated.

<img width="603" alt="Screenshot 2024-12-07 at 3 11 14 PM" src="https://github.com/user-attachments/assets/2798d12c-a00c-4587-a5a8-53660b511e90">

What this looks like graphically is that you have all your app components with their datamodels, and that is all represented as a Virtual DOM in JavaScript. Then after a UI interaction in a Browser like a button click, the new Virtual DOM is compared with old Virtual DOM and the differences are figured out. Those differences are then run through a React process called **reconciliation** where the actual Browser DOM updates based on the difference between the old and the new Virtual DOM. That step is really fast and it only happens once.

Both solutions the tightly coupled two-way data binding solution and the oneway data binding solution, work great on small apps where there are just a small number of UI elements. However, as apps get more complex, the tightly-coupled two-way data binding doesn't scale well. What happens is that many tightly-coupled model-view-controller loops must continually synchronize dragging down the performance of the app. That's essentially because it's slow to directly update the Browser DOM from JavaScript.

In React as you add more components, which is what happens as apps get more complex, there is no direct linkage to the faraway browser DOM. All those components when they need to update, simply update the local Virtual DOM stored in JavaScript, and that's lightening fast. Then after all the components update a single merge step is done and that step figures out what real DOM elements to change, and then in a single pass, JavaScript reaches out to the faraway DOM and does the update. Both conceptually and in the real world, this makes web apps written in React perform really well for their users.

# React Architecture
- React is 100% component based.
- All React Apps start by rendering a single component.
- Understand how that single component renders to a browser.

<img width="1226" alt="Screenshot 2024-12-07 at 3 58 51 PM" src="https://github.com/user-attachments/assets/8aea2418-e6a6-4e74-8946-6a1966fd0fb2">

Look at the example react function and look at the JavaScript return value. It specifically returns a new React element created by the library call **createElement**. Remember React with a capital R is the React library itself. It's likely that all the React apps you build will be using one of the popular tool chains. There is a good chance it will be Next.js as it's the most popular, but there are others also, including Remix, Gatsby, and RedwoodJS. If not you're likely using a custom webpack build that runs the **Babel Transpiler** directly allowing you to use JSX syntax in your apps.

We will be using Next.js which launches our App by instantiating the component on our behalf that is in the file ```/app/page.js```. This file contains the Root element of our React App. 

A **Library** is defined as a group of related functions or classes that perform unique operations to optmize well defined tasks.

React is essentially two libraries:
```js
import react from 'react';
import reactDOM from "react-dom"
```

One is ```react``` and the other one is ```react-dom```. The easiest way to see the source of each is to put the code from the [React site](https://react.dev/blog/2024/04/25/react-19-upgrade-guide#umd-builds-removed) that includes script tags and the SEM imports into HTML file.

```html
<script type="module">
  import React from "https://esm.sh/react@19/?dev"
  import ReactDOMClient from "https://esm.sh/react-dom@19/client?dev"
  ...
</script>
```

Let's include h1 tag just for reference.
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <script type="module">
        import React from "https://esm.sh/react@19/?dev"
        import ReactDOMClient from "https://esm.sh/react-dom@19/client?dev"
    </script>
</head>
<body>
    <h1>Hello from what is React</h1>
</body>
</html>
```

Open this in browser and look at the network tab. It shows amongst other things request for our react library and look at twhat was fetched. At the bottom you can see all the libarary calls enumerated. That includes not only low-level react calls like ```react.createElement``` but also the higher level ones, including all the ReactHooks like **useState** and **useEffect**.

Second for react-dom more of the same, that is ReactDOM is the Browser specific library that allows react to manipulate the browser DOM. It's totally upto you to decide how to use the calls in the library to build your apps.

A Framework on the otherhand is a collection of programming tools with a specific goal in mind like building a website or deploying a website. The typical requirement for a framework is that it is in control of the flow of the program you are creating unlike a library where you are in control.

If you are using just a library, the code you write calls the library code directly, and you choose how that code works. If you are using a Framework that framework itself calls your code and the framework contained calls to the library. Your program can also call the library directly. 

An example of a React Framework is **NextJS**. **NextJS** not only calls the react libraries it also uses Node, Webpack, SWC, which is the code transpiler written in Rust, they used to use Babel, PostCSS, Jest and many other technologies.

Bottomline, if you are building apps with React, you are likely to use a framework mostly because of all the time it will save you when building react apps for the web.