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

<img width="683" alt="Screenshot 2024-12-07 at 5 05 17 PM" src="https://github.com/user-attachments/assets/fba61390-82b5-4509-a1e1-7af3bfa1fda6">

Bottomline, if you are building apps with React, you are likely to use a framework mostly because of all the time it will save you when building react apps for the web.

# What It Means to be SPA
When a web page is requested by the browser, the page downloaded includes HTML and JavaScript and that from then on, there are **no more full-page requests** made from the browser.

Below is a very basic example of SPA:
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Singe Page App (SPA)</title>
    <script src="index.js"></script>
</head>
<body>
    <div id="root"></div>
</body>
</html>
```

index.js
```js
window.onload = () => {
  const rootElement = document.getElementById("root");
  rootElement.innerHTML = "Hello from my first SPA App!";
};
```

We need to make sure the page is fully rendered before executing any code that references elements on our HTML page. To do that we code the onload event, which is window.onload.
Technically it's a SPA because we use JavaScript after the page rendered to manipulate our UI, but we can make it little more interesting by adding a button that actually does something.

```js
window.onload = () => {
  const rootElement = document.getElementById("root");
  const button = document.createElement("button");
  button.innerHTML = "Click Me for current date";
  button.addEventListener("click", () => {
    button.innerHTML = new Date().toString();
  });
  rootElement.appendChild(button);
};
```

Now we know what a single page app is and we know that React is a library that we can make API calls to from our JavaScript that's running inside our browser.

Lets start out by modifying the simple SPA that we built above with just DOM calls to iterate over an array of numbers and then render each number:

```js
window.onload = () => {
  const rootElement = document.getElementById("root");
  const ints = [1, 2, 3];

  ints.forEach((i) => {
    let li = document.createElement("li");
    li.innerHTML = i;
    rootElement.appendChild(li);
  });
};
```

This is basic javascript updating the browser DOM. Let's do the same thing with React library.

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Simple React App No JSX</title>
    <script src="index.js"></script>
    <script type="module">
        import React from "https://esm.sh/react@19/?dev"
        import ReactDOMClient from "https://esm.sh/react-dom@19/client?dev"
      </script>
</head>
<body>
    <div id="root"></div>
</body>
</html>
```

Now let's replace index.js code using React. The First thing to do is to create our React virtual DOM root and assign it to a new variable, that's ```const root = ReactDOMClient.createRoot(root)```. From here on we will have all our interactions with react virtual DOM.

```js
window.onload = () => {
  const rootElement = document.getElementById("root");
  const root = ReactDOMClient.createRoot(rootElement);
  const ints = [1, 2, 3];
  const childrenElements = [];
  childrenElements.push(React.createElement("li", { key: ints[0] }, ints[0]));
  childrenElements.push(React.createElement("li", { key: ints[1] }, ints[1]));
  childrenElements.push(React.createElement("li", { key: ints[2] }, ints[2]));

  root.render(childrenElements);
};
```

Is this version faster. Not really but if the values associated with elements change React then figures out using the VirtualDOM what has changed and instead of having to replace the full DOM, can just surgically update what needs to be updated and replace just those items in the physical DOM.

# Creating a React App with the Next.js Toolchain/Framework
You can do everything that React does without a Toolchain and simply program the full React library by making React API calls directly, but almost nobody does that.

Unlike many other JavaScript frameworks out there, the core developers of React have tried very hard to focus on their mission, which is to build a JavaScript library that helps developer like us build high performance websites efficiently.

For a long time, they offered no toolchain to help us build our apps, but instead assumed we would create our own or that they would be provided by the community.

In 2016, those core developers at Facebook, finally decided that they needed to provide the community a standardized way to build React apps, and so they created the **Create React APP** project, and as is no surprise, it became very popular, and for a longtime, was the primary way most of us developers build apps with React.

**Create React APP** is no longer even mentioned in the Facebook docs, and the Opensource toolchain, **Next.js** is by far the most popular way to build React apps. It's actively maintained by **Vercel**, who also happens to have core developers from the React team on its payroll.

You'll still find lots of apps out there built with Create-React-APP, so it's certainly worth knowing but not recommended to start new app with it.

The Plan now is to build the identical app we already built with just the React API, but instead with a Toolcahin and JSX.

The way to create Next.js app is to type the command

```sh
npx create-next-app@latest 
```

Then it will ask series of questions
```
✔ What is your project named? … myapp
✔ Would you like to use TypeScript? … No
✔ Would you like to use ESLint? … Yes
✔ Would you like to use Tailwind CSS? … No
✔ Would you like your code inside a `src/` directory? … Yes
✔ Would you like to use App Router? (recommended) … Yes
✔ Would you like to use Turbopack for next dev? … No
✔ Would you like to customize the import alias (@/* by default)? … No
```

Select APP Router as yes, meaning that our routing will be served starting in the ```/src/app``` folder and specifically, ```page.js```. That means that when browsing to the base URL of the website, in our case in development mode that's ```localhost:3000```, the component that is default exported in the file ```src/app/page.js``` is rendered. In addition, if we created a new file in ```src/app/mylist/page.js``` then the URL at ```localhost:3000/mylist``` would render the default exported component in that file.

By default component in all of those special route files, ```page.js``` that is, is a server component. That means it will run inside your node server before rendering to the browser. This means React hooks like **useState** and **useEffect**, which are associated with the React page lifecycle running in the browser will not work.

Now to launch our app run below command:
```sh
npm run dev
```
There is lot of code generate in ```src/app``` folder but for our purposes, the only thing that we are currently interested in is the React root element associated with our app, that is the code in the file ```src/app/page.js```.

Let's remove all the code and replace it with just a simple functional component 

```js
import React from "react";

export default function () {
  return <div>Hello From Pluralsight</div>;
}
```

As mentioned before this component (src/app/page.js) is the server component, meaning ti is rendered in the node server. We can override this behaviour and **make it a pure spot component by adding a single line at the top 'use client' in quotes**.

```js
'use client'
import React from "react";

export default function () {
  return <div>Hello From Pluralsight</div>;
}
```

Let's remove the extra files in this folder we don't need. and add a simple global.css as below:
```css
html,
body {
  padding: 0;
  margin: 0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
}

a {
  color: inherit;
  text-decoration: none;
}

h1 {
  font-size: 2rem;
}

* {
  box-sizing: border-box;
}
```

What's important to take away from this so far is that:
- We never had to create a base HTML file and the JavaScript like we did with our simple React app without a toolchain.
- We don't have to create a Virtual DOM like we did previously.
- We don't have to explicitly render our Virtual DOM to the browser.

The good news is Next.js toolchain is taking care all of that for us. Let's create a simple list of numbers

```js
"use client";
import React from "react";

export default function () {
  return (
    <ul>
      <li>1</li>
      <li>2</li>
      <li>3</li>
    </ul>
  );
}
```

Now let's create a new Component which is just a JavaScript function list items. We could create in a spearate file but for demonstration we will create in the same file.

```js
"use client";
import React from "react";

function ListItems() {
  const ints = [1, 2, 3];
  return (
    <>
      {ints.map((id) => {
        return <li key={id}>{id}</li>;
      })}
    </>
  );
}

export default function () {
  return (
    <ul>
      <ListItems />
    </ul>
  );
}
```

# What F(G(X)) Means
If you have ever done any pure functional programming, a common expression people use to describe it is, it's the language of F(G(X)). What exactly does that mean?

Basically, it means that **functional programming is about Functional Composition**. Meaning that one function returns another function and that in turn return another function. It's an amazingly powerful way to build complex apps by combining simpler functions together.

It helps you build apps that contain more reusable code reducing duplication while also making them easier to test and find bugs in. This is how exaclty React builds component trees.

Looking at our current app we just built with just two functions, we could rename our main function to F, rename out ListItem function to G, and we essentially now have for our components F(G(X)), that is function F returns Function G. We don't have an X here because we haven't talked about React State but essentially X would be the React State passed in through the function call.

# What Makes React Apps Reactive
The react engine can maintain for every component in your app what we refer to as component state. State data is different than just data stored in JavaScript variables or objects in that when you update that state value, it persists and causes the React App to re-render itself with whatever the new value of state is. With React, you can't just declare a variable with let and expect that you can change its value and the page will re-render. The react engine has no idea of what's going on with variables you declare with let and would have no way to have the page UI update.

Let's now update our simple number rendering app to use component state, and that should help make clear how react works with state.

Let's move our array of numbers up in the component tree to the main component and pass that array into our list items component.

**The way we pass data from a parent to a child component is we create an attribute on a React element and then assign it to the JavaScript object or primitive we want to pass to the child components**

```js
"use client";
import React from "react";

function ListItems({ ints }) {
  return (
    <>
      {ints.map((id) => {
        return <li key={id}>{id}</li>;
      })}
    </>
  );
}

export default function () {
  const ints = [1, 2, 3];
  return (
    <ul>
      <ListItems ints={ints} />
    </ul>
  );
}
```

Now convert our ints array into React State. We know we are going to want a function that an event will call, so let's also add, while we are here at our top level component a new function definition named addValue. Have  The syntax is as below:
```js
"use client";
import React from "react";
import { useState } from "react";

function ListItems({ ints, addValue }) {
  return (
    <>
      <button onClick={addValue}>Add Item</button>
      {ints.map((id) => {
        return <li key={id}>{id}</li>;
      })}
    </>
  );
}

export default function () {
  const [ints, setInts] = useState([1, 2, 3]);
  function addValue() {
    const newVal = Math.max(...ints) + 1;
    setInts([...ints, newVal]);
  }
  return (
    <ul>
      <ListItems ints={ints} addValue={addValue} />
    </ul>
  );
}
```

**The way React works to pass values from a parent to a child component is to pass them as an attribute. The way child components can pass data to parent components is a little more indirect in that we pass a function to the child component and that child component can call that function to pass data back to the parent.**

You may be thinking that we really didn't pass any data up from our child to our parent component when the onClick event of the button fired. We can easily solve that by first adding an increment value parameter to our addValue function. Change the increment value from one to that new parameter, then instead of having our ListItem attribute pass the passed in addValue function back to our parent, we can replace that addValue function with an anonymous function that calls addValue.

```js
"use client";
import React from "react";
import { useState } from "react";

function ListItems({ ints, addValue }) {
  const increment = 3;
  return (
    <>
      <button onClick={() => addValue(increment)}>Add Item</button>
      {ints.map((id) => {
        return <li key={id}>{id}</li>;
      })}
    </>
  );
}

export default function () {
  const [ints, setInts] = useState([1, 2, 3]);
  function addValue(incrementValue) {
    const newVal = Math.max(...ints) + incrementValue;
    setInts([...ints, newVal]);
  }
  return (
    <ul>
      <ListItems ints={ints} addValue={addValue} />
    </ul>
  );
}
```

# How React Work With Browsers
React at it's core, is a user interface, UI, library. It's component based, meaning that the library itself manages the relationships between different UI components independent of how they are rendered, that is different devices have different methods and logic for how to render the actual UI presented to the user.

React separates the building and managing of components from their rendering to a device.

## Building Apps For React & React Native
- There is no "write once, run everywhere" for React and React Native.
- Separate Components required for UIs in React & React Native.
- You can build shared components between React & React Native.

React Native
```js
function App() {
    return (
        <View
            style={{
                fontSize: 20,
                fontWeight: "bold",
            }}>
            <Text>Hello From Pluralsight</Text>
        </View>
    )
}
```
React for Web
```js
function App() {
    return (
        <div>
            <b>Hello From Pluralsight!</b>
        </div>
    )
}
```

JSX syntax is identical but we use completely different component for Mobile & Web, however component structure & composition is identical for full apps, that is building component trees & passing data it's just normal JavaScript functions.

## Two libraries define React on the Web
- **React**
    - All about creating React Elements
    - Those elements have the ultimate purpose of creating UIs for a React app.
    - Typically involves linking components together in a nested component tree with a top element, the root element of the app.
    - React Libarary is the same whether you are making apps for web or mobile.
- **React DOM**
    - Just for building React apps that run in Web Browser. All of web browser have a DOM, a document object model.
    - React DOM is all about rendering elements to a web browser.
    - React DOM's primary purpose is to take a React app's root element and render it to a physical browser DOM.
    - React DOM is all about the what and the where to render. 

## Understanding React reconciliation
At the start of every React App, whether you are using a framework or not, is a call to ```ReactDOM.createRoot``` that returns an object that is the root of your App. After that when you call ```root.render``` the React APP renders to the web browser. 

index.js
```js
import ReactDOM from "react-dom";
const container = document.getElementById('root');
const root = ReactDOM.createRoot(container);

const RootComponent = () => <div>Hello From PluralSight!</div>
root.render(
  <RootComponent />
)
```

Assuming that you have rendered UI that is interactive, meaning there are events in your app that update component state, everytime this happens, React on your behalf essentially calls render again and re-renders your app.

In the todo app above as soon as we enter new item and click add item button, we get a new list item added to our unordered list. What happened for the third list item to be rendered?
React had to use both the original virtual DOM previously saved as well as the new Virtual DOM created after the user pressed the Add Item button.

<img width="1108" alt="Screenshot 2024-12-08 at 9 51 01 PM" src="https://github.com/user-attachments/assets/bd0eb519-0ae2-4a1c-bd95-eb15ea89d3f4">

React uses these two copies of Virtual DOMs to do a step called **Reconciliation**. This means that the old DOM is compared to the new DOM and only the elements that change are updated. For our case, it's very simple. The new Virtual DOM has just one more element in it than the previously stored old Virtual DOM. From the compare, the necessary DOM calls are created such that just one node can be added to the DOM and no full DOM re-render is needed.

## Complex React APP Reconciliation
In more complex APPs for example the Facebook website itself, you can imagine lots of elements can change and also the shape of the componenets can change. The challenge is to efficiently figure out what minimal diffs need to be applied to the browser to make the UI correct for the new updated DOM to be reflected.

If you have n components or React Elements in your App's component tree that the complexity of calculating what components need to be updated is O(n^3). That means if you have 200 components in your app, you will potentially need to do 6 million comparisons to figure out the appropriate differences between your old and new DOM.

Luckily there are lots of clever shortcuts the React team has figured out to make that comparison really fast and nowhere near to the O(n^3) calculations. Things like knowing what the React element type is can make a huge difference. That is all the todo items are the same type, where as the TodoItemList component is different type.

Bottomline, the reconciliation step is very fast in real apps because of some awesome job React team has done in optimizing this step.

# Distributed Components
We want to introduce Server Components without calling them Server Components. Basically, think about React apps that have both components running in a node server and components running in a client browser working together, of course, as distributed components in a React App.

Currently, only viable distributed React Apps are using React Server & Client Components.

Since the node server is running basically the same javascript engine as is running in the browser, it can run React Components and then pass data as receiving parameters down into the browser where client compoents then run. 

<img width="763" alt="Screenshot 2024-12-08 at 10 13 12 PM" src="https://github.com/user-attachments/assets/9d53d568-ce9d-48cb-ae96-94c0a51d19e3">

In future, we will likely see React Components running inside a separate process in the browser. That process could be running inside a **Web Worker**, it might be running with **Web Assembly**, also called **WASM** and it could be using a **Service Worker**, or for that matter, something not even invented yet.

For now though, the only implementation available for running distributed React is using **Server Components** and **Client Components** with those server components running a node.

## What's Behind Server Component Technology 
Using just ```index.html``` and pure javascript file with no react, we served a webpage to a browser with a simple numbers list. It did it by generating 100% of the list inside the javascript and after the webpage loaded, that JavaScript wrote that list directly to the browser DOM. Our point back then was to demonstrate a pure SPA.

Later we converted it to React App, also a SPA and added a button that when clicked, added more numbers to the list. Let's take a step backwards and create a pure javascript app and add to it a button with a simple event handler that adds numbers to the list when clicked.

```html
<!DOCTYPE html>
<html lang="en">
<body>
    <div id="root"></div>
    <script>
        const rootElement = document.getElementById("root");
        const numbers = [1, 2, 3, 4, 5];
        
        const generateListHTML = (numbers) => `
            <ul>
                ${numbers.map((number) => `<li>${number}</li>`).join("")}    
            </ul>
        `;
        
        const generateButtonHTML = () => `
            <button id="addItemButton">Add Item</button>
        `;
        rootElement.innerHTML = `
            ${generateListHTML(numbers)}${generateButtonHTML()}
        `;
        const addItemButton = document.getElementById("addItemButton");
        const incrementValue = 3;
        const addItem = () => {
            const newNumber = numbers[numbers.length - 1] + incrementValue;
            numbers.push(newNumber);
            rootElement.querySelector("ul")
                .insertAdjacentHTML("beforeend", `<li>${newNumber}</li>`)
        };
        addItemButton.addEventListener("click", addItem);
    </script>
</body>
</html>
```

Now we have single HTML file that has exact same capability. As this HTML page has nothing to do with any particular server, that is it doesn't require a node server to run, we can simply invoke VS Code plugin LiveServer.

We did this exercise to understand how a React App that is served from a node server and includes React Server Components can first render a static HTML page directly to the browser, and then subsequently, run in SPA mode processing browser events that is like button clicks and update the DOM to reflect changes or new numbers being added to the list.

Let's evelove this JavaScript to replicate the rendering scenario we just described and then we will do the exact same thing in React and it will help understand what React is doing under the cover for us.

```js
const http = require("http");

const server = http.createServer(async (req, res) => {
  res.setHeader("Content-Type", "text/html");
  res.writeHead(200);
  res.end(`
    <!DOCTYPE html>
        <html lang="en">
        <body>
            <div id="root"></div>
            <script>
                const rootElement = document.getElementById("root");
                const numbers = [1, 2, 3, 4, 5];
                const generateListHTML = (numbers) =>  \`
                    <ul>
                        \${numbers.map((number) => \`<li>\${number}</li>\`).join("")}    
                    </ul>
                \`;
                const generateButtonHTML = () => \`
                    <button id="addItemButton">Add Item</button>
                \`;
                rootElement.innerHTML = \`
                    \${generateListHTML(numbers)}\${generateButtonHTML()}
                \`;
                const addItemButton = document.getElementById("addItemButton");
                const incrementValue = 3;
                const addItem = () => {
                    const newNumber = numbers[numbers.length - 1] + incrementValue;
                    numbers.push(newNumber);
                    rootElement.querySelector("ul")
                        .insertAdjacentHTML("beforeend", \`<li>\${newNumber}</li>\`)
                };
                addItemButton.addEventListener("click", addItem);
            </script>
        </body>
        </html>
    `);
});

server.listen(3000, "127.0.0.1", () => {
  console.log("Server running at http://127.0.0.1:3000/");
});
```

Now run node server as:
```sh
node server.js
```

Now we can hit http://127.0.0.1:3000 at it serves the same html. So now same compiler is running in node server instead of in browser, we can refactor some of this JavaScript passed down to the browser and bring it back to the server.

```js
const http = require("http");

const numbers = [1, 2, 3, 4, 5];
const generateListHTML = (numbers) => `
    <ul>
        ${numbers.map((number) => `<li>${number}</li>`).join("")}    
    </ul>
`;
const generateButtonHTML = () => `
    <button id="addItemButton">Add Item</button>
`;

const server = http.createServer(async (req, res) => {
  const initialListHTML = generateListHTML(numbers);
  const buttonHTML = generateButtonHTML();
  res.setHeader("Content-Type", "text/html");
  res.writeHead(200);
  res.end(`
    <!DOCTYPE html>
        <html lang="en">
        <body>
            <div id="root">${initialListHTML}${buttonHTML}</div>
            <script>
              const numbers = ${JSON.stringify(numbers)};
                const rootElement = document.getElementById("root");
                const addItemButton = document.getElementById("addItemButton");
                const incrementValue = 3;
                const addItem = () => {
                    const newNumber = numbers[numbers.length - 1] + incrementValue;
                    numbers.push(newNumber);
                    rootElement.querySelector("ul")
                        .insertAdjacentHTML("beforeend", \`<li>\${newNumber}</li>\`)
                };
                addItemButton.addEventListener("click", addItem);
            </script>
        </body>
        </html>
    `);
});

server.listen(3000, "127.0.0.1", () => {
  console.log("Server running at http://127.0.0.1:3000/");
});
```

Look how we have moved generateListHTML and generateButtonHTML out as a variable and used as a variable. This means that when request is made from the browser to this node server, the first thing that happens before anything is downloaded to the local browser is that the HTML to be rendered is figured out. Then in the code downloaded to the browser, specifically inside the root div, these variables are inserted.

This means on the First page downloaded to the browser, instead of just empty div rendering with id root, that div now contains the full list of initial numbers, as well as the HTML for the button.

Next to support our AddItem button being clicked, we do need to know the initial array of numbers in our Browser JavaScript, so an easy way to get it is to create a new const named numbers and then json encode the javascript array, so that the JavaScript code that gets downloaded to the browser looks like a normal javascript array declaration.

Restart the node server and look at the page source. Element with id root has already rendered unordered list in HTML. Since view-source looks at the downloaded HTML before any javascript runs you can be sure that browser user will immediately see the numbers list as fast as the page comes up. No waiting for JavaScript execution like we had with the pure single page app. Also notice the const numbers declaration in the Script tag.

After clicking the addItem button we can see the items being added but if we look at the page-source we still see the 5 items but if we use chrome Debug Tool and inspect the page we can see all the dynamically added elements.

## Convert a Server Rendered App is JS to React
Above we developed code to show the basics of how to server render HTML and still use JavaScript to enhance it's interactivity. This works well for trivial example, but as Apps get more complex trying to keep track of what's running where gets really complicated with just JavaScript & DOM calls.

Lets look at the exact same scenario but solve using React Server & Client component. **React Server & Client Components are 100% a React thing and not a toolchain like next.js thing**. 

However, toolchain is involved in settings the basics like where the root element of your app is running and whether that element is a server or a client component. As we are using **next.js**, the rule when using **App Router** is the root element is always defined in a file named ```/app/page.js```, it could also be jsx, ts or tsx.

Below are server components

page.js
```js
import React from "react";
import AddItemButton from "./add-item-button";
import SharedDataProvider from "./shared-data-provider";
import NumbersList from "./numbers-list";

const numbers = [1, 2, 3, 4, 5];

export default async function Home() {
  return (
    <div className="container">
      <SharedDataProvider initialLastNumber={numbers[numbers.length - 1]}>
        <NumbersList numbers={numbers} />
        <AddItemButton increment={3} />
      </SharedDataProvider>
    </div>
  );
}

function ListItems({ ints, addValue }) {
  const increment = 3;
  return (
    <>
      <button onClick={() => addValue(increment)}>Add Item</button>
      {ints.map((id) => {
        return <li key={id}>{id}</li>;
      })}
    </>
  );
}
```

numbers-list.js
```js
import NewNumbers from "./new-numbers";

export default async function NumbersList({ numbers }) {
  return (
    <ul>
      {numbers.map((number) => (
        <li key={number}>{number}</li>
      ))}
      <NewNumbers />
    </ul>
  );
}
```

By default it's server component otherwise we mention 'use client' for **Client Component**.

Below are client components:

new-numbers.js
```js
"use client";

import { useSharedData } from "./shared-data-provider";

export default function NewNumbers() {
  const { newNumbers } = useSharedData();

  return (
    <>
      {newNumbers.map((number) => (
        <li key={number}>{number}</li>
      ))}
    </>
  );
}
```

shared-data-provider.js
```js
"use client";

import React, { createContext, useContext, useState } from "react";

const SharedDataContext = createContext();

export function useSharedData() {
  const contextValue = useContext(SharedDataContext);
  if (!contextValue) {
    throw new Error(
      "useSharedData must be used within a " + "SharedDataProvider"
    );
  }
  return contextValue;
}

export default function SharedDataProvider({ initialLastNumber, children }) {
  const [newNumbers, setNewNumbers] = useState([]);

  function addNewNumber(incrementValue) {
    const lastNumber = newNumbers[newNumbers.length - 1] || initialLastNumber;
    setNewNumbers([...newNumbers, lastNumber + incrementValue]);
  }

  return (
    <SharedDataContext.Provider
      value={{
        newNumbers,
        addNewNumber,
      }}
    >
      {children}
    </SharedDataContext.Provider>
  );
}
```