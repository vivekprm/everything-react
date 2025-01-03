# What Problem Does React Solve?

- Static UIs are easily manageable.
- Dynamic Changes are hard
- Lots of code needed that's difficult to maintain.

React helps, it brings structure with components.

# Core React Features

- Offers structure by letting us compose the UI with separate reusable components. Components are the building blocks of a React application and can have a changeable state.
- The UI of a component is declared in JavaScript and that's unconventional. But we'll see that this has a huge advantage. It gives React the ability to make it easier to dynamically change the UI of components when the state changes. The declared UI will be rendered to a browser or other output automatically by React.
- React is very efficient because when the UI updates, it will not just bluntly replace the whole rendered component. Instead, a smart mechanism called **reconciliation** will figure out what actually changed and it will only update that.

# Components & JSX

- Components are JavaScript functions that return JSX, a syntax that seems like HTML.
- JSX looks like HTML but it is not, it's just an alternative way to write JavaScript.
- JSX stands for JavaScript Extension. It extends the JavaScript syntax to make it easier for us developers to write UIs in an intuitive hierachical way.
- Browser doesn't understand JSX. JSX has to be transformed by a tool to JavaScript. The typical tool for that is called **Babel**. We can go to [Babel website](https://babeljs.io/repl#?browsers=defaults%2C%20not%20ie%2011%2C%20not%20ie_mob%2011&build=&builtIns=false&corejs=3.21&spec=false&loose=false&code_lz=KYDwDg9gTgLgBAE2AMwIYFcA29noHYDGMAlhHnABIQC2wAFAJQDeAUHHFMDOlOXW-zgAeAIJgwAPiEB6MZIEMWAXyA&debug=false&forceAllTransforms=false&modules=false&shippedProposals=false&evaluate=false&fileSize=false&timeTravel=false&sourceType=module&lineWrap=true&presets=env%2Creact%2Cstage-2&prettier=false&targets=&version=7.26.4&externalPlugins=&assumptions=%7B%7D) and try out how it converts JSX to JavaScript.

## JSX to Markup

- So from JSX, we can get to **createElements** statements using Babel tool.

```
<h1>This is a banner</h1> ----> React.createElement("h1", null, "This is a banner")
```

- What happens next and how does the browser know what to render. Each time, browser has to render a component a JavaScript library called React DOM uses to create element code to generate the actual HTML elements for the browser.

## Separation of Concerns

- React separates its core library from the library that is responsible for rendering components. The latter is called react-dom for browser applications but react can also be used for mobile applications using React Native.
- In case of react native the core library is still used but react-dom will be different library that can render mobile controls.

Let's look at another component:

```js
const Greeting = () => (
  <div>
    <Banner />
    <h2 className="highlight">Greetings!</h2>
  </div>
);
```

In this case, JSX returned has a div that contains the Banner component from the previous slide, it is rendered in this component using the same HTML like element syntax we used with the div and the h2. That **div and h2, in fact, are also components**. They're just built-in components that correspond to DOM elements. Since we are targetting browser here the components used will result in the rendering of HTML elements with the same name. So div for example, will result in the rendering of a div in the browser.

Now let's say you are not using a browser as a rendering target like we do in this course, but the mobile device for instance, in that case you won't use the same built-in components but the ones that corresponds to controls that are available on a mobile device.

The casing of the names of DOM related components built into React is camelcase.
Custom components created by us are Pascal Cased where the first letter is capital letter.

# Application Structure

Create project using next.js

```sh
npx create-next-app globomantics
```

# Modules

What are modules? Well they are not React specific, they are part of JavaScript.

Lets say we have `module.js` file with a function called doSomething.

```js
const doSomething = () => {
    ...
}
export {doSomething};
```

To make this file a module, we need to export something from it using the export keyword. In this case we are exporting the function. Exporting it makes it possible for other modules to import and use it. Should the file contain other things that are not exported, these things are just not available to other modules.

You can also export multiple members of the module by just separating them with commas. Members of the module can be classes, functions, object basically anything that can exist in a variable. Any module that wants to import things from this module does so using the import keyword followed by what it wants to import between curly braces.

anotherModule.js

```js
import { doSomething } from "./module";
doSomething();
```

It's also possible for a module to export one default member using the default keyword.

module.js

```js
const doSomething = () => {
    ...
}
export default doSomething;
```

Importing side then doesn't have to use the curly brace syntax and also it can use any name.

anotherModule.js

```js
import do from "./module";
do();
```

It's also possible to have combination of default export and normal export.

## Reasons to Use Modules

First of all they bring structure in the code base. Just like components are the building blocks of the UI, modules are the building blocks in the code.

- They make easy to reuse code across the application without using global variables.
- They bring Encapsulation. Everything that is not exported from a module remains private for the module.
- They are needed for bundling. The **bundler** tool will try to create one big JavaScript file out of all these seperate little files. Without modules, it would have no idea about the order of the individual files.

# Running the Finished Project

https://github.com/RolandGuijt/ps-react-fundamentals

# Adding New Components

Lets remove everything that we had generated for next.js project and add a directory called components. Add banner and app components.

banner.js

```js
const Banner = () => {
  return (
    <header>
      <div>
        <img src="./GloboLogo.png" alt="logo" />
      </div>
      <div>Providing houses all over the world</div>
    </header>
  );
};
export default Banner;
```

app.js

```js
import Banner from "./banner";

const App = () => {
  return <Banner />;
};
export default App;
```

Now add app component in index.js

```js
import App from "../components/app";

const Index = () => <App />;
export default Index;
```

# Detecting Problems & Debugging

Part of the `Next.js` tooling is **ESLint**, which is able to detect problems in the code, and it also helps with code styling.
`Next.js` comes with a ruleset for React preinstalled and you can run the tool that check these rules by typing `npm run lint` in the terminal.

When we debug we see breakpoint is hit twice and tht's because React strict mode is on. This will try to anticipate possible problems with the application and for that it executes function components twice. It won't happen in production build but if you don't like it while developing, you can turn it off in the `next.js` configuration file.

How come we see JSX code while debugging in browser? Afterall we learned that browser doesn't understand JSX.
That's because of the tooling. It creates source map files during the dev build that link the generated Babel code to the original source.

When debugging in chrome it's worthwile to install a plugin called "React Developer Tools" This will get us two extra tabs:

- Profiler
- Components

# Styling Components

There are couple of ways to style components.

- Next.js specific
  - In \_app.js CSS files can be loaded that can be applied throughout the application.
- We deleted \_app.js because we are going to load the CSS in the HTML root document.
- It's also possible to have a CSS file with styles that are isolated for just one Component.
- Use **style** attribute on a component.

Most web applications use a standard CSS file to get started style-wise.

In the demo application we are using bootstrap mainly to position the components. A typical way to load CSS file is in the HTML root document, the file with the Html & the Head tags.

But in `next.js` that's nowhere to be found in the project template. Should you use create-react-app, you will find it. It's simply a file called `index.html`. But in `next.js` it's provided behind the scenes. It's possible to access it though. Just create an \_document.js

This is the default content for it:

```js
import { Html, Head, Main, NextScript } from "next/document";

export default function Document() {
  return (
    <Html lang="en">
      <Head />
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  );
}
```

Now we can add a link pointing to bootstrap as below:

```js
import { Html, Head, Main, NextScript } from "next/document";
import Link from "next/link";

export default function Document() {
  return (
    <Html lang="en">
      <Head>
        <link
          href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css"
          rel="stylesheet"
          integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH"
          crossOrigin="anonymous"
        />
      </Head>
      <body>
        <div className="container">
          <Main />
        </div>
        <NextScript />
      </body>
    </Html>
  );
}
```

## Bootstrap Grid System

The screen is divided in 12 equally wide invisible columns. When the CSS class **row** is applied to the component a new row will be created in the grid for the children of the component. In the banner component we can assing css className row to the header component.

Two divs are are the children. They can now claim a certain width of the row using a **col** class. `col-5` means 5 columns wide. so first div accupies 5 columns and second div 7 columns.

To leave some space between the row and what comes next, we can apply a second css class to header `mb-4` means margin bottom 4. So below is the complete banner.js

```js
const Banner = () => {
  return (
    <header className="row mb-4">
      <div className="col-5">
        <img src="./GloboLogo.png" alt="logo" className="logo" />
      </div>
      <div className="col-7 mt-5">Providing houses all over the world</div>
    </header>
  );
};
export default Banner;
```

## CSS Modules

- CSS modules make it possible to import a css file in a component.
- That CSS will be only available to that component as opposed to a CSS file like bootstrap that we loaded in the HTML root document. That is available for whole application for all the components.
- CSS modules are not a react feature. The functionality is provided by **webpack**, the tool that bundles CSS files for Next.js and create react app as well as others.
- With CSS modules, we can have a nicely separated CSS per component. That means when styling for a component has to be changed, there is no need to open a giant CSS file to find the classes used.
- There is no need to worry that we unintentionally change the styling of another component that uses the same class we're changing.
- But we learned all the CSS files are bundled by the tooling, so how does this work?
  - Tooling also takes care of giving all class names a unique name, and it of course will automatically apply the new names in the components.

To give the Banner component it's own css create a file called `banner.module.css`

```css
.logo {
  height: 150px;
  cursor: pointer;
}
```

Now to use this inside banner component.

```js
import styles from "./banner.module.css";

const Banner = () => {
  return (
    <header className="row mb-4">
      <div className="col-5">
        <img src="./GloboLogo.png" alt="logo" className={styles.logo} />
      </div>
      <div className="col-7 mt-5">Providing houses all over the world</div>
    </header>
  );
};
export default Banner;
```

## Style Attribute

- Just like HTML elements, components also have a style attribute.
- Difference is that in React, styles take an object that contains CSS style properties. So an expression is need to use it.
- Object can be an inline object or an existing object.

To style the banner text, we can set the style attribute on the surrounding div.

```js
import styles from "./banner.module.css";

const subtitleStyle = {
  fontStyle: "italic",
  fontSize: "x-large",
  color: "coral",
};
const Banner = () => {
  return (
    <header className="row mb-4">
      <div className="col-5">
        <img src="./GloboLogo.png" alt="logo" className={styles.logo} />
      </div>
      <div className="col-7 mt-5" style={subtitleStyle}>
        Providing houses all over the world
      </div>
    </header>
  );
};
export default Banner;
```

Use of the style attribute is however discouraged, for the same reason style attribute in HTML is discouraged.

CSS should be in separate CSS files to keep it apart from the components.
Doing it like that is also better for performance.

# Hooks, Props & State

Components can receive arguments from other components using props. Prop values can be set by other components using HTML like attribute syntax.

```jsx
<Component myprop="somevalue" />
```

Here a prop called **myprop** is set to the string value **somevalue**. But prop values aren't limited to strings, they can be anything you can put in a variable, such as objects, arrays and functions.

The component receiving the props can access them by using an object that is passed into the components function.

```jsx
const Components = (props) =>
```

We will revisit the banner component and display the header text different everytime.

**Props are read-only, a component should never modify its own props**. One of the reasons is many props are passed on to the other components by reference using objects or arrays for example. Let's say the props is an object because the reference to the object is passed on when a component changes the prop, the object will change for every component that has a reference. These other components have no way to detect that it was changed. To avoid problems around that, in React we are using specific pattern. **Prop data flow goes one way, components receiving props are not allowed to change them.**.

banner.js

```js
import styles from "./banner.module.css";

const subtitleStyle = {
  fontStyle: "italic",
  fontSize: "x-large",
  color: "coral",
};
const Banner = ({ headerText }) => {
  return (
    <header className="row mb-4">
      <div className="col-5">
        <img src="./GloboLogo.png" alt="logo" className={styles.logo} />
      </div>
      <div className="col-7 mt-5" style={subtitleStyle}>
        {headerText}
      </div>
    </header>
  );
};
export default Banner;
```

app.js

```js
import Banner from "./banner";

const App = () => {
  return <Banner headerText="Providing houses all over the world." />;
};
export default App;
```

## Children Prop

When using the banner it might feel more natural to, instead of using an explicit prop, just put in the text as a child of the Banner component as below:

```js
<Banner>Providing houses all over the world</Banner>
```

That's no problem React supports that. **Props** object contains a special property called **children** which contains all the markup present between the start and end tag. In this case **children** is string with text.

aap.js

```js
import Banner from "./banner";

const App = () => {
  return <Banner>Providing houses all over the world.</Banner>;
};
export default App;
```

banner.js

```js
import styles from "./banner.module.css";

const subtitleStyle = {
  fontStyle: "italic",
  fontSize: "x-large",
  color: "coral",
};
const Banner = ({ children }) => {
  return (
    <header className="row mb-4">
      <div className="col-5">
        <img src="./GloboLogo.png" alt="logo" className={styles.logo} />
      </div>
      <div className="col-7 mt-5" style={subtitleStyle}>
        {children}
      </div>
    </header>
  );
};
export default Banner;
```

**children** of a component can be more complex than a string.

## Fragments & Mapping Data to JSX

Let's introduce new Component that renders the list of houses. Add file `houseList.js`

```js
import React from "react";

const HouseList = () => {
  return (
    <div>
      <div className="row mb-2">
        <h5 className="themeFontColor text-center">
          Houses currently on the market
        </h5>
      </div>
      <table className="table table-hover">
        <thead>
          <tr>
            <th>Address</th>
            <th>Country</th>
            <th>Asking Price</th>
          </tr>
        </thead>
        <tbody></tbody>
      </table>
    </div>
  );
};

export default HouseList;
```

In this JSX the first div is not needed, it's just there to satisfy the requirement that only on parent can be there inside JSX. For such cases we can use `React.Fragment` or shorthand syntax `<>`.

```js
import React from "react";

const HouseList = () => {
  return (
    <React.Fragment>
      <div className="row mb-2">
        <h5 className="themeFontColor text-center">
          Houses currently on the market
        </h5>
      </div>
      <table className="table table-hover">
        <thead>
          <tr>
            <th>Address</th>
            <th>Country</th>
            <th>Asking Price</th>
          </tr>
        </thead>
        <tbody></tbody>
      </table>
    </React.Fragment>
  );
};

export default HouseList;
```

themeFontColor css class is present in `global.css`, rest of the classes are bootstrap classes. If you look at the component body of the table is empty. Data for the houses will come from an array.

`map` is a javascript function available on array, it takes each item in the array and map it to something else, producing a new array. E.g.

```js
const numbers = ["one", "two", "four"];
const numbersPrefixed = numbers.map((n) => "Number " + n);
```

houseList.js

```js
import React from "react";

const houses = [
  {
    id: 1,
    address: "12 Valley of Kings, Geneva",
    country: "Switzerland",
    price: 900000,
  },
  {
    id: 2,
    address: "89 Road of Forks, Bern",
    country: "Switzerland",
    price: 500000,
  },
];

const HouseList = () => {
  return (
    <>
      <div className="row mb-2">
        <h5 className="themeFontColor text-center">
          Houses currently on the market
        </h5>
      </div>
      <table className="table table-hover">
        <thead>
          <tr>
            <th>Address</th>
            <th>Country</th>
            <th>Asking Price</th>
          </tr>
        </thead>
        <tbody>
          {houses.map((h) => (
            <tr key={h.id}>
              <td>{h.address}</td>
              <td>{h.country}</td>
              <td>{h.price}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );
};

export default HouseList;
```

## key prop

You might notice key attribute for each row, it's needed wherever arrays of React elements are created.
Why it's needed?
Let's say we had two houses:

```js
{
  address: "432 Tree Lane";
}
{
  address: "495 Newton St";
}
```

What would you expect React to do if a house is added to the array as below?

```js
{
  address: "33 Palm Dr";
}
{
  address: "432 Tree Lane";
}
{
  address: "495 Newton St";
}
```

The most efficient would be that the rendered HouseList will remain intact and that only the new item was inserted. But rightnow, there's no way React can link the generated JSX to the item in the array, so it simply doesn't know how to upgrade the list of rendered houses. The only option is to just refresh the whole list. This might not be the much of the performance impact with just 3 items in the array but what it the list grows to hundred or thousands of items. We need something to help React link the array item to the rendered house. That something is special key prop, value of which is used internally by React.

With the helpp of this key React knows that the two existing items were already rendered, and it can now just add a new one. If you don't have an id property, you are free to comeup with anything unique (may be combination of properties). If there is nothing to uniquely identify an item, as a last resort, you can use the item index that is provided by the map function. But this can cause problem when the order of the items change, so only use them when there are no other options.

## Extracting Components

Rightnow houses column has only 3 columns so the JSX is manageable. It would however become a different story if a few more columns would be added. We could create a separate component for the house row to keep the JSX readable and manageable.

houseRow.js

```js
import React from "react";

export const HouseRow = ({ address, country, price }) => {
  return (
    <tr>
      <td>{address}</td>
      <td>{country}</td>
      <td>{price}</td>
    </tr>
  );
};
```

In this case, HouseList can pass each individual prop as:

houseList.js

```js
import React from "react";
import { HouseRow } from "./houseRow";

const houses = [
  {
    id: 1,
    address: "12 Valley of Kings, Geneva",
    country: "Switzerland",
    price: 900000,
  },
  {
    id: 2,
    address: "89 Road of Forks, Bern",
    country: "Switzerland",
    price: 500000,
  },
];

const HouseList = () => {
  return (
    <>
      <div className="row mb-2">
        <h5 className="themeFontColor text-center">
          Houses currently on the market
        </h5>
      </div>
      <table className="table table-hover">
        <thead>
          <tr>
            <th>Address</th>
            <th>Country</th>
            <th>Asking Price</th>
          </tr>
        </thead>
        <tbody>
          {houses.map((h) => (
            <HouseRow
              key={h.id}
              address={h.address}
              country={h.country}
              price={h.price}
            />
          ))}
        </tbody>
      </table>
    </>
  );
};
export default HouseList;
```

That works but house can have multiple properties and it can be tedius to pass each property as props individually. Instead we can pass it using spread operator like below.

```js
<tbody>
  {houses.map((h) => (
    <HouseRow key={h.id} {...h} />
  ))}
</tbody>
```

However be careful with this. Let's say we added new property to house that the HouseRow doesn't need. Then each property will be passed on as individual prop anyway, and that can negatively impact performance. So passing h as house object is the better.

Finally let's format the price. For that create a directory called helpers and add `currencyFormatter.js`. It uses JavaScripts internationalization feature.

```js
const currencyFormatter = Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
  maximumFractionDigits: 0,
});

export default currencyFormatter;
```

And modify the HouseRow component to use it:

```js
export const HouseRow = ({ house }) => {
  return (
    <tr>
      <td>{house.address}</td>
      <td>{house.country}</td>
      <td>{currencyFormatter.format(house.price)}</td>
    </tr>
  );
};
```

## Hooks

Hooks play a very important role when working with function components.

- A hook is a function.
- Name of the function always has **use** prefix.
- It's purpose is to encapsulate complexity.
  - By simply calling a function we can use pretty complicated functionality without the need to exactly know how it works.
- Built-in Hooks make it possible to use React's internal features within function components.
- In addition to React's built-in hooks it is also possible to create your own hooks.

### Rules of Hooks

- Hooks should only be called at the top level.
  - It's forbidden to call them conditionally, wrapping them with if statement for example.
  - By applying this rule, it is ensured that they will be called in the same order everytime the component's function is executed. React needs that to make hook function correctly.
- Only call hooks in function components.
  - Call from outside the function component will result in error.
  - The only exception is custom hook. So a custom hook may call other hooks.

### State Hook

**Props** are used to pass data into a component. Props are set by a parent component.
**State** is data that is kept internally by a component. But why do we need state?

As we learned, in React we don't write HTML, unlike many other libraries and frameworks that are about UIs.

```json
[
  {
    "id": 1,
    "address": "12 Valley of Kings, Geneva",
    "country": "Switzerland",
    "price": 900000
  },
  {
    "id": 2,
    "address": "89 Road of Forks, Bern",
    "country": "Switzerland",
    "price": 500000
  }
]
```

When we have to render data like this array, we write JSX as an instruction on how to render it.

```js
{
  houses.map((h) => <HouseRow key={h.id} house={h} />);
}
```

Ans react uses the JSX to generate HTML. So how can we add row to the table?
We are not writing HTML, so changing the table itself is off the table. Changing the JSX won't help because it is fine as it is. The only way to do it, is to add a row to the houses array. It's basically one way data flow. We change the data upstream and the table that is rendered downstream is ultimately just a reflection of it.

But if we add an item to the array, there is no way that React can know the item was added, so the UI won't update and that's not we want. We need state to fix that. So let's modify houseList.js to use state.

```js
import React, { useState } from "react";
import { HouseRow } from "./houseRow";

const houseArray = [
  {
    id: 1,
    address: "12 Valley of Kings, Geneva",
    country: "Switzerland",
    price: 900000,
  },
  {
    id: 2,
    address: "89 Road of Forks, Bern",
    country: "Switzerland",
    price: 500000,
  },
];

const HouseList = () => {
  const [houses, setHouses] = useState(houseArray);
  return (
    <>
      <div className="row mb-2">
        <h5 className="themeFontColor text-center">
          Houses currently on the market
        </h5>
      </div>
      <table className="table table-hover">
        <thead>
          <tr>
            <th>Address</th>
            <th>Country</th>
            <th>Asking Price</th>
          </tr>
        </thead>
        <tbody>
          {houses.map((h) => (
            <HouseRow key={h.id} house={h} />
          ))}
        </tbody>
      </table>
    </>
  );
};

export default HouseList;
```

By calling the set function in this case setHouses, react will know and it will re-render the UI automatically if needed.

#### Setting State

We are changing the state of the HouseList component by adding a new House object to the array.

```js
import React, { useState } from "react";
import { HouseRow } from "./houseRow";

const houseArray = [
  {
    id: 1,
    address: "12 Valley of Kings, Geneva",
    country: "Switzerland",
    price: 900000,
  },
  {
    id: 2,
    address: "89 Road of Forks, Bern",
    country: "Switzerland",
    price: 500000,
  },
];

const HouseList = () => {
  const [houses, setHouses] = useState(houseArray);
  const addHouse = () => {
    setHouses([
      ...houses,
      {
        id: 3,
        address: "32 Valley Way, New York",
        country: "USA",
        price: 1000000,
      },
    ]);
  };
  return (
    <>
      <div className="row mb-2">
        <h5 className="themeFontColor text-center">
          Houses currently on the market
        </h5>
      </div>
      <table className="table table-hover">
        <thead>
          <tr>
            <th>Address</th>
            <th>Country</th>
            <th>Asking Price</th>
          </tr>
        </thead>
        <tbody>
          {houses.map((h) => (
            <HouseRow key={h.id} house={h} />
          ))}
        </tbody>
      </table>
      <button className="btn btn-primary" onClick={addHouse}>
        Add
      </button>
    </>
  );
};

export default HouseList;
```

We are free to use state hook multiple times in one component.

```js
const HouseList = () => {
  const [houses, setHouses] = useState(houseArray);
  const [counter, setCounter] = useState(0);
  setCounter(counter + 1)
  ....
}
```

counter here is primitive type in javascript not a refrence type. With premitive type, react will compare the previous value to the new value. Only if there is a difference, it will update the UI, provided that we use counter somewhere in the JSX.

By the way **set** function also can take a function as parameter that gets the current state.

```js
const [counter, setCounter] = useState(0);
setCounter((current) => counter + 1);
```

current here is equivalent to the value of counter most of the time, but when multiple calls to the set function are done, React batches the calls for efficiency, only updating the state value when the batch completes. Using this pattern it's guaranteed that current contains the value that was set by the previous call to the set function.

### Props & State Interaction

Props are read only and they shouldn't be changed by the component, so you are probably surprised to hear that "prop value can change". How does this work?
What is a prop for one component is often state for another. Let's consider **HouseRow** component created earlier. For **HouseRow** house is a prop but for **HouseList** the house instance that is passed to **HouseRow** is part of it's state. What happens if **HouseList** does a call to **setHouses**, passing in a new house array instance, like we do in the **addHouse** function

# Component Rendering & Side Effects

Re-rendering means the component's function is executed again. It can be because of the state change or rendering of parent component.

But that doesn't mean that entire component is refreshed in the Browser. React uses **Reconciliation** where the browser will only update the parts that were actually changed.

React stores state internally by using in-memory arrays. So when a component uses two calls to useState, first one for the houses array and second one for the counter, it will keep array containing two elements. The first one contains the houses array and second one the counter value.
That's the reason useState hook can't be called conditionally. Doing this will mess up this internal system of state keeping.

## The Rendering Process

Our demo application so far has this component hierarchy:

We have an App component that renders the Banner & HouseList components and the HouseList component renders a number of HouseRow components. When the state of HouseList changes, it will re-render. That doesn't necessarily happen in that same millisecond. **React will first just flag the component and for performance reasons, it will just queue a render**. Then after a period of time, which appears instantaneous for humans, the actual rerendering will happen.

React will walkthrough the tree of components, finding all components that need to be rendered. In this case it's only one i.e. HouseList. It will not only render the Flag component but also all the child components. Basically everything under HouseList component tree.

It can have cascading effect of rerenders, is that a bad thing?
Not necessarily. Remember reconciliation, the things that eat the most performance are the actual browser updates, and React minimizes those, but it's something to keep in mind when you design your application, if thousands of functions are fired in response to the change of one piece of state anywhere in your application, you can imagine that is not too good for performance.

## Pure Functions & Memoizing Component

React relies on Pure Functions to perform efficiently. **A pure function is a function that always returns the same result**. e.g.

```js
const returnNumber = () => 42;
```

Add function is also a pure function because as long as same value of a and b are used it returns the same result.

```js
const add = (a, b) => a + b;
```

Pure functions are:

- Easy to test
- Predicatable
- Reliable
- Cacheable

All these benefits are what React needs to do it's rendering work efficiently. So **a component's function should be a pure function**. In component terms, that means that **given the same prop values and the same state, the function should always return the same JSX**.

Despite that we saw that all the HouseRow components are re-rendered anyway when the state of it's parent, HouseList changes. But there is a way to memoize, cache the output for a component so that it doesn't re-render when it's prop value remains the same. It's done by wrapping the component with `React.memo`.

Now, in addition to or instead of the **HouseRow** we already had, we can export the Memoized version.

```js
import currencyFormatter from "@/helpers/currencyFormatter";
import React from "react";

const HouseRow = ({ house }) => {
  return (
    <tr>
      <td>{house.address}</td>
      <td>{house.country}</td>
      <td>{currencyFormatter.format(house.price)}</td>
    </tr>
  );
};

const HouseRowMem = React.memo(HouseRow);
export default HouseRow;
export { HouseRowMem };
```

Now if we have a break point in **HouseRow** and click Add button, breakpoint is hit only once. So only the new row is rendered, since the prop value for other **HouseRow** remained the same.

Maybe you'll have tendency now to wrap every component with **React.memo**. However **React.memo** has its overhead in terms of performance, and React's rerendering cycle is highly optimized. Therefore you should only use it when:

- It's faster.
- When you can measure it's faster.
  - You can do that using profile tool that comes with React Development Plugin.
- In general it's faster only when it's a pure functional component you're wrapping and when it renders often with the same prop values.
- JSX returned shouldn't be trivial, like our HouseRow component.
  - Because it's so small and React is so optimized that performance gain in this case is insignificant.

Beware that **React.memo** will only shallowly compare complex objects when they are passed in as props.

https://react.dev/reference/react/memo

## Side Effects & Effect Hook

Pure functions might sound straight forward enough. But in practice we have to do things in components that are not so predictable and reliable. And that's okay as long as you keep in mind that these operations should be set aside. They should not be part of the pure function. Such operations are called **Side Effects**.

Basically whenever we reach out to something that is not within the realm of React, we have to use an effect because the results are unpredictable and may be unreliable.

Examples of Effect are:

- API Interaction
- Use Browser APIs (e.g. document, window)
- Using timing functions (e.g. setTimeout)

### Effect Hook

To perform an effect in the function of a component, the effect hook is used.

```js
useEffect(() => {
  // perform the effect
});
```

**useEffect** takes a function as a parameter. This function will be executed automatically after React is done running the component's pure function and the browser has been updated.

In the function, the effect is performed. So we can fetch data from an API here for example. So we can get rid of the houseArray in HouseList component and initialize the houseArray state with empty array.

fetch is an asynchrnous operation it returns a promise. So ideally we want to put await in front of fetch and capture resulting response object into a variable. But to make this work, we have to make the function passed into useEffect to aync. But this will make function return a promise and useEffect can't work with that.

Instead we can wrap the call to fetch in extra function that is async and call that.

```js
useEffect(() => {
  const fetchHouses = async () => {
    const response = await fetch("/api/houses");
    const houses = await response.json();
    setHouses(houses);
  };
  fetchHouses();
});
```

Let's add the api in pages/api the js files that you see doesn't run in browser, they run on the server using nodejs.

Developing client-side code together with server-side code like this is pecific to next.js. But we don't want complexity of writing an api and talking to database etc. Instead add a json file called houses.json, which is read and manipulated by API code.

Implementing an API like this isn't something you should do in production. It's just for demo purposes.

Problem is when we press Add button the extra row appears just briefly, let's return to the code to investigate what's going on. When the house list is initially rendered houses state is set to an empty array. When useEffect is called, the function contained in the useEffect will not be called immediately. First rest of the pure function executes, so the JSX is returned without any HouseRows because the houses array is empty.

Now React will run the effect, so call goes to API, result is read and we call setHouses to change the houses state. Now the component re-renders and function will be re-executed. So we start at the first line again because it's re-render the initial value of the state is ignored, instead existing state is used, the houses we fetched earlier through API. Then again useEffect containing function doesn't fire yet. We first return the JSX, which will now contain the HouseRows and browser is updated to show them.
Now the effect fires again fetching the houses, changing the state causing re-render, the state now contains the newly fetched houses which is the same list as before. SO the JSX returned is exactly the same and React determines that the browser doesn't have to be updated so it remanins static in the browser. After that effect function is called again fetching the houses, changing the state causing the rerender. We are stuck in infinite loop.

The solution to this is to instruct react to only run the useEffect function in certain cases. That's done by specifying a dependency array.

```js
const [counter, setCounter] = useState(0);
useEffect(() => {
  document.title = counter;
});
```

Do we need the effect to run everytime a component rerenders?
No, just when the component is initially rendered and when the counter changes. To make that happen we pass the dependency array as the second parameter to use effect as below:

```js
const [counter, setCounter] = useState(0);
useEffect(() => {
  document.title = counter;
}, [counter]);
```

But in the case of the effect in HouseList, there is no dependency. All we want is that effect function is executed only once, just when the component initially renders. To make that happen we can specify an empty dependency array.

```js
useEffect(() => {
  const fetchHouses = async () => {
    const response = await fetch("/api/houses");
    const houses = await response.json();
    setHouses(houses);
  };
  fetchHouses();
}, []);
```

What if there are multiple effects in a component that have to be executed.

```js
const [counter, setCounter] = useState(0);
useEffect(() => {
  document.title = counter;
}, [counter]);

useEffect(() => {
  // fetch from API
}, []);
```

In that case don't try to squeeze them into one call to useEffect. Multiple calls are supported, and they can each have their own dependency array.

And if needed you can also return a function from the effect to clean things up.

```js
useEffect(() => {
  // subscribe
  return () => {
    // unsubscribe
  };
}, []);
```

You can also return a function from the effect to clean things up.
What if you, for example, are subscribing to an event stream from some API? You will want to unsubscribe when the component is unmounted. That means removed from the UI. You can do that in this function. But beware that this function is not only called when a component is removed, it is also called everytime before the effect function fires again. In this case that won't happen. But if you have dependencies in the dependency array or no dependency array at all, keep this in mind.

### Memo Hook

It could be handy to optimize the performance of your components. We have already seen that component output can be memoized but values inside components can be too.

Lets say a calculation has to be done involving a list of houses that is quite time-consuming

```js
const result = timeConsumingCalculation(houses);
```

We could put this line of code into our component but now the calculation will be done on every re-render as we learned re-render can occur frequently, this will slow down the application.

To gain performance, we can memoize the value that is returned using the memo hook.

```js
const result = useMemo(() => {
  return timeConsumingCalculation(houses);
}, [houses]);
```

The parameters of the useMemo hook are very much like the ones from useEffect. The first parameter is a function that does the calculation. The calculation will occur when the component is first rendered and when houses changes because houses is in the dependency array, that is the second parameter.
If the component is re-rendered without a change to houses `useMemo` will simply return the value that was calculated previously without running the function.

So Is it a good idea to just wrap any operation in a useMemo call? Doesn't hurt right?
Well, it does because of the overhead of the hook. Again you should measure if this is really faster before putting it in.

### Ref Hook

The Ref hook can be used to store values that are persisted between renders. That sounds familiar because that's what state does. **The difference is that modifying a ref value doesn't cause a re-render**. If for example you want to count number of time HouseRow component has rendered, we can't use counter state and call setCounter and useEffect because that will trigger a re-render.causing an infinite loop.

Instead we can call useRef hook with an initial value of 0. useRef returns an object that has a current property that contains the current value. This value can be modified directly and it won't cause re-render.

```js
const [houses, setHouses] = useState([]);
const counter = useRef(0);
useEffect(() => {
  const fetchHouses = async () => {
    const response = await fetch("/api/houses");
    const houses = await response.json();
    setHouses(houses);
  };
  fetchHouses();
  counter.current++;
}, []);
```

When a reference type is passed to useRef the ref hook guarantees that the same reference is returned in the current property across re-renders. May be that's why hook is called Ref.

It's also often used to gain access to JavaScript DOM objects and components

```js
const TextInputWithFocusButton = () => {
  const inputEl = useRef(null);
  const onButtonClick = () => inputEl.current.focus();
  return (
    <>
      <input ref={inputEl} type="text" />
      <button onClick={onButtonClick}>Focus The Input</button>
    </>
  );
};
```

in JSX inputEl is overwritten with the input DOM element object. We can access it by reading the current property of the ref. In this case, focus is called on the input element, but all JavaScript members available on DOM objects are available here. This object is the same object you get when you make call using `document.getElementById` for example, in any javascript application.

# Conditional Rendering & Shared State

You can conditionally apply almost anything in JSX. Let's say we want to conditionally apply

```js
const HouseRow = ({ house }) => {
  let priceTd;
  if (house.price < 50000)
    priceTd = <td>{currencyFormatter.format(house.price)}</td>;
  else
    priceTd = (
      <td className="text-primary">{currencyFormatter.format(house.price)}</td>
    );
  return (
    <tr>
      <td>{house.address}</td>
      <td>{house.country}</td>
      {priceTd}
    </tr>
  );
};
```

But in this case we can save some lines of code using an expression, something like:

```js
<td className={`${house.price >= 500000 ? "text-primary" : ""}`}>
  {currencyFormatter.format(house.price)}
</td>
```

We can also render price td only when price is present:

```js
{
  house.price && (
    <td className={`${house.price >= 500000 ? "text-primary" : ""}`}>
      {currencyFormatter.format(house.price)}
    </td>
  );
}
```

Lets add House.js to render house details like longer description and picture when we click on any row in HouseList.

When a row is clicked the list is replaced by another component. We can make that happen by rendering the HouseList and a new component called House conditionally.

We will add House component in house.js

```js
import defaultPhoto from "@/helpers/defaultPhoto";
import React from "react";

export const House = ({ house }) => {
  return (
    <div className="row">
      <div className="col-6">
        <div className="row">
          <img
            className="img-fluid"
            src={
              house.photo ? `./houseImages/${house.photo}.jpeg` : defaultPhoto
            }
            alt="House pic"
          />
        </div>
      </div>
    </div>
  );
};
```

It uses house prop to display all data of an individual house. There is nothing really new in this component.
photo property on the house that is returned from the API is a string. If there is a photo defined the value will be truthy. In that case using the photo property, we look for the JPEG image in the houseImages directory. We have added houseImages directory with the images to public. When the photo property isn't truthy something called defaultPhoto is used instead. Which we have added in the helpers directory as:

```js
const defaultPhoto = "data:image/jpeg;base64,base64";

export default defaultPhoto;
```

For a public facing production app you'll probably want to cache the photo. Return a physical image instead.

It still need a parent component that can render it. Which one will be suitable?
We want to replace HouseList with House somehow and HouseList is rendered in the App root component and that would be the place to do it.
Here we render either House or HouseList based on state value.

```js
const [selectedHouse, setSelectedHouse] = useState();
```

Here we are not passing anything in useState, that means selectedHouse initially will be undefined.

If selectedHouse is truthy we render House else HouseList. selecteHouseRow state should change when a HouseRow is clicked in HouseList. The App Component can't do it. It has no way of knowing that the HouseRow which is a couple of component levels down is clicked. So capturing that event somehow in App in not an option.

What we can do is give HouseRow a way to set the state for App. As you know, only way to change state is to call the set function.

Here is what we did in schema :

<img width="1030" alt="Screenshot 2024-12-11 at 9 23 46 PM" src="https://github.com/user-attachments/assets/958f233d-1819-422c-ad30-a7331a704428" />

On App we defined state, which provided the setSelectedHouse function. We passed it as prop to the HouseList. And HouseList renders HouseRow components that get the function too via prop. The function is passed by reference. So these are not copies of the function, they all refer to the same function.

Now when a HouseRow is clicked, the function is called. By doing so, the state of the App component is changed, and all children re-render. The selectedHouse state no contains a House, so instead of HouseList, House is rendered.

Because the whole application re-renders, when the state in the root component changes, it is not wise to have too much of it there. Just have state in the root component that doesn't fit anywhere else. The state we have here is a good example since this is state that controls the overall composition of components.

**To prevent unnecessary re-renders and making the application overly complicated, only share state between components when needed**. Place state as low as you can go in the component hierarchy.

app.js

```js
const App = () => {
  const [selectedHouse, setSelectedHouse] = useState();
  return (
    <>
      <Banner>Providing houses all over the world.</Banner>
      {selectedHouse ? (
        <House house={selectedHouse} />
      ) : (
        <HouseList selectHouse={setSelectedHouse} />
      )}
    </>
  );
};
export default App;
```

houseList.js

```js
...
<tbody>
  {houses.map((h) => (
    <HouseRow key={h.id} house={h} selectHouse={selectHouse} />
  ))}
</tbody>
...
```

houseRow.js

```js
const HouseRow = ({ house, selectHouse }) => {
  return (
    <tr onClick={() => selectHouse(house)}>
      <td>{house.address}</td>
      <td>{house.country}</td>
      <td>{currencyFormatter.format(house.price)}</td>
    </tr>
  );
};
```

## Mounting & Unmounting

When components appear in browser DOM, they are mounted, their state is initialized and effects run. When the application starts, all components for the initial composition are mounted. When HouseRow is clicked, HouseList and it's rows are removed from the DOM. They are unmounted. It's important to realize that with that all state of these components is destroyed. That sounds rather permanent.
But at some point, the user of the application may decide to return to the list of houses. We haven't implemented that functionality yet but when we do, HouseList and it's rows will be mounted again. Since all memory of the previous HouseList and HouseRows is destroyed, the components will be initialized as if they never existed before. So the state initializes with the initial value, and the effects will run again.

So for HouseList, that means that houses will be refetched from the API. To have more control around what child components put into the state, a wrapper can be used.

## Function Wrappers & The Callback Hook

HouseList and HouseRow get the setSelectedHouse function via a prop, and that is the function that is coming from the call to the state hook in App. Everything works perfectly as it is now, but we're basically giving full control of the selected house state to these child components.

May be they will accidentally won't put in a house object, but another object that is not related or some string or number. When this application gets larger, this might lead to serious bugs. To remedy that, we can create a wrapper function that accepts a house objects and does the call to **setSelectedHouse** and instead of passing setSelectedHouse to HouseList we now pass wrapper to it.

```js
const App = () => {
  const [selectedHouse, setSelectedHouse] = useState();
  const setSelectedHouseWrapper = (house) => {
    setSelectedHouse(house);
  };
  return (
    <>
      <Banner>Providing houses all over the world.</Banner>
      {selectedHouse ? (
        <House house={selectedHouse} />
      ) : (
        <HouseList selectHouse={setSelectedHouseWrapper} />
      )}
    </>
  );
};
```

The advantage of this approachis that the App component now still is in full controlof its own state. The **setSelectedHouse** function remains encapsulated in the component. We can now add checks to the wrapper function to make sure the thing that is passed into the function is really a house.

There is something to keep in mind when creating functions like this though.

- On each re-render the function object is created, creating a new reference.
  - That's not a problem unless the function is passed into a component that is memoized with React.memo. Because the object reference changes, the components will re-render maybe unintentionally.
  - Another example is **when the function is used in the dependency array of an effect, for example, the effect's function will execute**. To prevent that, **the callback hook (i.e. useCallback) can be used**. This will preserve the same function reference across re-renders, unless something in it's dependency array changes, it memoizes the containing function object.
    - It comes with some overhead, so you should only use it when it's really necessary.

```js
const [selectedHouse, setSelectedHouse] = useState();
const setSelectedHouseWrapper = useCallback((house) => {
  setSelectedHouse(house);
}, []);
```

In this case, the containing function will only be created when the component is mounted. For every re-render, the reference remains the same. SetState functions, such as our setSelectedHouse, never have to be wrapped with **useCallback** because React will make sure they are not recreated on every render.

## Delegating State to a Custom Hook

Right now **HouseList** is fetching houses data, and it is displaying it. Introducing a custom hook is good way to separate these two concerns. Create a directory **hooks** and add **useHouses.js**. Remember hooks should always have the **use** prefix in their name. A hook is just a function like a component, which is exported by a module. The difference is that hook doesn't return JSX, but it can use other hooks. hook function can return the houses state.

useHouses.js

```js
import { useEffect, useState } from "react";

const useHouses = () => {
  const [houses, setHouses] = useState([]);
  useEffect(() => {
    const fetchHouses = async () => {
      const response = await fetch("/api/houses");
      const houses = await response.json();
      setHouses(houses);
    };
    fetchHouses();
  }, []);
  return { houses, setHouses };
};

export default useHouses;
```

So let's modify the houseList.js

```js
import React from "react";
import HouseRow from "./houseRow";
import useHouses from "@/hooks/useHouses";

const HouseList = ({ selectHouse }) => {
  const { houses, setHouses } = useHouses();
  const addHouse = () => {
    setHouses([
      ...houses,
      {
        id: 3,
        address: "32 Valley Way, New York",
        country: "USA",
        price: 1000000,
      },
    ]);
  };
  return (
    <>
      <div className="row mb-2">
        <h5 className="themeFontColor text-center">
          Houses currently on the market
        </h5>
      </div>
      <table className="table table-hover">
        <thead>
          <tr>
            <th>Address</th>
            <th>Country</th>
            <th>Asking Price</th>
          </tr>
        </thead>
        <tbody>
          {houses.map((h) => (
            <HouseRow key={h.id} house={h} selectHouse={selectHouse} />
          ))}
        </tbody>
      </table>
      <button className="btn btn-primary" onClick={addHouse}>
        Add
      </button>
    </>
  );
};

export default HouseList;
```

Of-course you are also free to create a wrapper function for setHouses in the hook.

\*\*A custom hook is a function which can accept any parameters you want. It can also return anything you want. When you return a hook's state the component that uses the custom hook will re-render when the state changes. This gives us React developers a great amount of flexibility. We can separate different parts of component functionality to prevent huge component function.

Since custom hook is separate from the component itself we can reuse it. But keep in mind that when a custom hook is reused the state for each call to it is isolated. That means **if useHouses will be used in another component that would get it's own house state that is separate from the houses state in HouseList.**.

In the case of **useHouses** that would mean that for the other component, the houses will be fetched from the API again. But we will get into **a way to share state globally with a feature called context**.

### Adding Additional State of a Custom Hook

We are making our custom hook a bit more interesting and the user interface a bit more user-friendly. Rightnow, there is no indication that the application is loading something from the API. The user just sees an empty table for a second. To change that add `loadingStatus.js` file in helpers directory, which exports a simple object with properties that contain all the possible loading statuses.

```js
const loadingStatus = {
  loaded: "loaded",
  isLoading: "Loading...",
  hasErrored: "An error occured while loading",
};

export default loadingStatus;
```

Let's also add a new Component, **LoadingIndicatore**, it takes loadingState prop and displays it.

```js
import React from "react";

export const LoadingIndicator = ({ loadingState }) => {
  return <h3>{loadingState}</h3>;
};
```

Now in the useHouses hook introduce a new state called loadingState, which we keep upto date while loading.

```js
import loadingStatus from "@/helpers/loadingStatus";
import { useEffect, useState } from "react";

const useHouses = () => {
  const [houses, setHouses] = useState([]);
  const [loadingState, setLoadingState] = useState(loadingStatus.isLoading);
  useEffect(() => {
    const fetchHouses = async () => {
      try {
        setLoadingState(loadingStatus.isLoading);
        const response = await fetch("/api/houses");
        const houses = await response.json();
        setHouses(houses);
        setLoadingState(loadingStatus.loaded);
      } catch {
        setLoadingState(loadingStatus.hasErrored);
      }
    };
    fetchHouses();
  }, []);
  return { houses, setHouses, loadingState };
};

export default useHouses;
```

Now in houseList.js we can export loadingState and use it.

```js
import React from "react";
import HouseRow from "./houseRow";
import useHouses from "@/hooks/useHouses";
import loadingStatus from "@/helpers/loadingStatus";
import { LoadingIndicator } from "./loadingIndicator";

const HouseList = ({ selectHouse }) => {
  const { houses, setHouses, loadingState } = useHouses();

  if (loadingState != loadingStatus.loaded) {
    return <LoadingIndicator loadingState={loadingState} />;
  }
  const addHouse = () => {
    setHouses([
      ...houses,
      {
        id: 3,
        address: "32 Valley Way, New York",
        country: "USA",
        price: 1000000,
      },
    ]);
  };
  return (
    <>
      <div className="row mb-2">
        <h5 className="themeFontColor text-center">
          Houses currently on the market
        </h5>
      </div>
      <table className="table table-hover">
        <thead>
          <tr>
            <th>Address</th>
            <th>Country</th>
            <th>Asking Price</th>
          </tr>
        </thead>
        <tbody>
          {houses.map((h) => (
            <HouseRow key={h.id} house={h} selectHouse={selectHouse} />
          ))}
        </tbody>
      </table>
      <button className="btn btn-primary" onClick={addHouse}>
        Add
      </button>
    </>
  );
};

export default HouseList;
```

It is also possible to return null in place of LoadingInticator, in that case it will hide itself from rendering.

### Increasing the Reusability of Custom Hook

In near future other components will need to fetch from the API using other URLs. This loading functionality is now fixed to fetching just list of houses. It would be great if we could somehow reuse this loading functionality, for example for a new component called Bids, which we will introduce later. TO facilitate that, we're adding a more generic hook called **useGetRequest**. It takes a URL for the request as a parameter

```js
import loadingStatus from "@/helpers/loadingStatus";
import { useState } from "react";

const useGetRequest = (url) => {
  const [loadingState, setLoadingState] = useState(loadingStatus.isLoading);

  const get = async () => {
    setLoadingState(loadingStatus.isLoading);
    try {
      const rsp = await fetch(url);
      const result = await rsp.json();
      setLoadingState(loadingStatus.loaded);
      return result;
    } catch {
      setLoadingState(loadingStatus.hasErrored);
    }
  };
  return { get, loadingState };
};

export default useGetRequest;
```

Now in useHouses hook we can remove loadingState and modify it as:

```js
const useHouses = () => {
  const [houses, setHouses] = useState([]);
  const { get, loadingState } = useGetRequest("/api/houses");
  useEffect(() => {
    const fetchHouses = async () => {
      const houses = await get();
      setHouses(houses);
    };
    fetchHouses();
  }, [get]);
  return { houses, setHouses, loadingState };
};
```

We see a warning, because we are using get function in useEffect, we should include it in dependency array. It also mentions the option to remove dependency array altogether, but that's not an option for us because that would cause useEffect to fire on every re-render, so lt's put in get in dependency array.

When we try things out, it doesn't seem to work correctly. There is a constant re-render going on while the loading indicator stays on and the API is queried continuously. Let's try to figure it ut step by step.

HouseList calls useHouses, which will trigger the effect in that hook, we wait for the Get request to finish and update the state but now multiple HouseList rerenders have happened. First the loadingState changes from loading to loaded and then the houses state is updated. When the HouseList re-renders it will call the **useHouses** again and useHouses call **useGetRequest**. Because **useGetRequest** is called again, the **get** function is recreated. Remember a function is just an object and we are getting a new reference here. But the **get** funtion is in dependency array of **useEffect** and **useHouses**.

The function in **useEffect** is called again, which causes re-renders where the **get** function is re-recreated again, causing a re-render. In short, another in-finite loop.
We can solve this by using the **useCallback** hook with the **get** function with a dependency array that has th URL. It has to be in there because it's an external dependency.

```js
import loadingStatus from "../helpers/loadingStatus";
import { useCallback, useState } from "react";

const useGetRequest = (url) => {
  const [loadingState, setLoadingState] = useState(loadingStatus.isLoading);

  const get = useCallback(async () => {
    setLoadingState(loadingStatus.isLoading);
    try {
      const rsp = await fetch(url);
      const result = await rsp.json();
      setLoadingState(loadingStatus.loaded);
      return result;
    } catch {
      setLoadingState(loadingStatus.hasErrored);
    }
  }, [url]);
  return { get, loadingState };
};

export default useGetRequest;
```

This will ensure that **get** function object isn't re-created but the same reference is used across re-renders unless the URL changes. But since every call to useGetRequest is a request to one particular URL, it's okay.

You could argue if the **useHouses** hook is still needed?
You could put it's **useEffect** call in the **HouseList** component again and remove a level of abstraction. It's upto us.

# Context & Navigation

Using **context** we can pass data through the component tree without having to pass props down manually to every level. It's suitable for functionality like navigation.

Some state in an application is relevant for a large number of components. The state in demo application that manages the different component compositions is an example. From the user's perspective, it is used to navigate through the application.

<img width="1030" alt="Screenshot 2024-12-11 at 9 23 46 PM" src="https://github.com/user-attachments/assets/958f233d-1819-422c-ad30-a7331a704428" />

Rightnow, the set function with which we can do the navigating is passed on from component to component. But what if, in the future, after implementing a couple of new features the component hierarchy looks more like this (see below pic).

In this situation the set function must be passed in via a prop to every component that needs to navigate. It works, but passing it in manually to every component level is cumbersome. It would be great if state could be defined somewhere so that all components could read it without the need to pass it on with props to every component that needs it. React provides the possibility to do just that by using context.

## Declaring Context

Here's an overview of how context works:

- Context has to be created. It is given default value as parameter.

```js
const context = React.createContext("default value");
```

- Context object has to be made available for the component that provides the context to its children as well as the child components that want to access the value provided by the context.
- In the component that wants to pass context to its children, a component called **Provider** contained in the **context** object can be rendered, providing a certain value. If no value is given default value given in step 1 is used.

```js
<context.Provider value="some value">// children</context.Provider>
```

- The **Provider** component doesn't have a visual appearance. It's job is to just provide the value to all child components.
- **Children** can now optionally read the context value. It doesn't matter if it's a direct child or a child that is further down in the hierarchy.

```js
const value = useContext(context);
```

- The value is read using the **useContext** hook, which takes the context object as parameter.
- In JSX there is also the possibility to use the **Consumer** component that is also present on the context. In an expression, a function is defined that gets the context value, which we can use to render something that uses it.

```js
<context.Consumer>
  {value => /* render something based on value */}
</context.Consumer>
```

**Important**: When the **context** value changes all children that access the context value will re-render.

In demo application having state in App that sets a selectedHouse was a great idea but it isn't sufficient if the application expands in functionality because there will be components that aren't about houses at all.
The selected house state only works to switch between the **HouseList** and **House** components. We could create more state to support other components but the logic that determines what has to be displayed quickly will become messy and unmaintainable.

So we are creating a more generic solution to do navigation. Create a file called `navValues.js` in helpers directory that contains all the different components compositions for our application.
At this point there are just two, **home** where **HouseList** is rendered and **house** for the **House** details.

```js
const navValues = {
  home: "Home",
  house: "House",
};
export default navValues;
```

In the App component we will get rid of selectedHouse state.

```js
import React, { useState } from "react";
import Banner from "./banner";
import HouseList from "./houseList";
import { House } from "./house";
import navValues from "@/helpers/navValues";

const navigationContext = React.createContext(navValues.home);
const App = () => {
  return (
    <navigationContext.Provider value={navValues.home}>
      <Banner>Providing houses all over the world.</Banner>
    </navigationContext.Provider>
  );
};
export { navigationContext };
export default App;
```

Context is only useful if the value can be changed. So we are introducing a **nav** state and instead of providing static values use state.

```js
const navigationContext = React.createContext(navValues.home);
const App = () => {
  const [nav, setNav] = useState(navValues.home);
  return (
    <navigationContext.Provider value={nav}>
      <Banner>Providing houses all over the world.</Banner>
    </navigationContext.Provider>
  );
};
export { navigationContext };
export default App;
```

Now let's create a wrapper function.

```js
const App = () => {
  const navigate = useCallback((navTo) => setNav(navTo), []);
  const [nav, setNav] = useState(navValues.home);
  return (
    <navigationContext.Provider value={nav}>
      <Banner>Providing houses all over the world.</Banner>
    </navigationContext.Provider>
  );
};
export { navigationContext };
export default App;
```

Now you would be inclined to pass this navigate function to the children by passing in an object to the value of the provider that contains both, the state and the function.

```js
return (
  <navigationContext.Provider value={{ nav, navigate }}>
    <Banner>Providing houses all over the world.</Banner>
  </navigationContext.Provider>
);
```

Be careful, with this however, since this object will be recreated with each rerender and thus all children consuming the value will re-render too. We want to leave the re-render responsibility to the state itself. So it's better to include the function in the state. Instead of just the nav value, an object is passed in with current property containing the navValue and the function and we have to adjust the navigate function too.

```js
const App = () => {
  const navigate = useCallback(
    (navTo) => setNav({ current: navTo, navigate }),
    []
  );
  const [nav, setNav] = useState((current) => navValues.home);
  return (
    <navigationContext.Provider value={nav}>
      <Banner>Providing houses all over the world.</Banner>
    </navigationContext.Provider>
  );
};
export { navigationContext };
export default App;
```

Now a way is needed to figure out what should be rendered. We could write a switch statement using an expression below banner, but let's keep this JSX nice and clean and use a component **ComponentPicker**.

**ComponentPicker** is passed the current nav location with a prop. ComponentPicker have a switch statment.

```js
import navValues from "@/helpers/navValues";
import React from "react";
import HouseList from "./houseList";
import { House } from "./house";

const componentPicker = ({ currentNavLocation }) => {
  switch (currentNavLocation) {
    case navValues.home:
      return <HouseList />;
    case navValues.house:
      return <House />;
    default:
      return (
        <h3>No component for navigation value {currentNavLocation} found</h3>
      );
  }
};

export default componentPicker;
```

House component needs selected house as prop. Now we will see how we can use context in HouseRow component.

```js
import currencyFormatter from "@/helpers/currencyFormatter";
import React, { useContext } from "react";
import { navigationContext } from "./app";
import navValues from "@/helpers/navValues";

const HouseRow = ({ house }) => {
  const { navigate } = useContext(navigationContext);
  return (
    <tr onClick={() => navigate(navValues.house, house)}>
      <td>{house.address}</td>
      <td>{house.country}</td>
      <td>{currencyFormatter.format(house.price)}</td>
    </tr>
  );
};

const HouseRowMem = React.memo(HouseRow);
export default HouseRow;
export { HouseRowMem };
```

Notice we are passing second parameter house to navigate function. Let's add support for that in App component.

```js
const App = () => {
  const navigate = useCallback(
    (navTo, param) => setNav({ current: navTo, param, navigate }),
    []
  );
  const [nav, setNav] = useState((current) => navValues.home);
  return (
    <navigationContext.Provider value={nav}>
      <Banner>Providing houses all over the world.</Banner>
    </navigationContext.Provider>
  );
};
export { navigationContext };
export default App;
```

Now in House component also we can use useContext as below:

```js
import defaultPhoto from "@/helpers/defaultPhoto";
import React, { useContext } from "react";
import { navigationContext } from "./app";

export const House = () => {
  const { param: house } = useContext(navigationContext);
  return (
    <div className="row">
      <div className="col-6">
        <div className="row">
          <img
            className="img-fluid"
            src={
              house.photo ? `./houseImages/${house.photo}.jpeg` : defaultPhoto
            }
            alt="House pic"
          />
        </div>
      </div>
    </div>
  );
};
```

Now everything works but it would be nice if we could return to HouseList somehow may be by clicking on the logo for example. Since all the child components has access to context, Banner component also can take advantage of it.

## Navigation Libraries

Our navigation logic may be sufficient for this application but what if more features are needed? For example when we navigate right now, the URL in the browser stays as it is. May be user would expect to return to the previous composition if they press the back button in the browser and a deep link to a particular house could be great so the other sites can link to it for SEO. We could build this and maybe more ourselves, but there is also the option to go with the routing system built into **next.js** if you use that or use an extra JavaScript library. **React Router** is the most popular option.

**Note**: React Router doesn't work with **next.js** because **next.js** uses server-side features that **React Router** doesn't support. So it's popularity comes from users from other frameworks like **Create React App**. It provides a number of components to setup routes with.
E.g.

```js
<Router>
  <HousesContext.Provider value={allHouses}>
    <div className="container">
      <Header subtitle="Providing houses all over the world" />
      <HouseFilter allHouses={allHouses} />
      <Routes>
        <Route
          path="/searchresults/:country"
          element={<SearchResults />}
        ></Route>
        <Route path="/house/:id" element={<HouseFromQuery />}></Route>
        <Route
          path="/"
          element={<FeaturedHouse house={featureHouse} />}
        ></Route>
      </Routes>
    </div>
  </HousesContext.Provider>
</Router>
```

**Routes** are basically the link between a URL and a component composition. To navigate to such a URL, a navigation function or a special **Link** component can be used.

If you need the extra features it's recommended to use third party navigation libraries such as **ReactRouter**. As long as you consider the extra dependency on the library and realize that all the features that these types of libraries bring are built around what you've learned in state using context.

## When and where should you use **Context**?

In the example application, context is used in the root component but it can be introduced at any level in the component hierarchy, for example at the level of **HouseList**, putting in a provider for a certain value there. Then that value will be available for all **HouseList** children.

So when it should be used?
Well bascially, if you have to pass the same state to many components via props. However, you should be aware of the implications of using context.

- The main one is that potentially many components will re-render when it changes.
- Another implication is that the component consuming it will rely on a context being present and that makes component reuse more difficult.
- And also realize that state provided by the context is basically hidden state. When you start writing anew component, you have no idea it's there unless you havea through understanding of all the rst of the source code.

# User Input & Forms

The input elements accept user inputs, under the hood this input is stored in memory. So these already have a state, it's just not react state.
When the form is submitted in a JavaScript Application that doesn't use React, we would have to:

- Write code that gets us a reference to the input as an object
- And extract the current value before it can, for example be posted to an API.
- React applications rely on state managed by React, so we have to convert the internal state handling of the elements to the React state.
- We have to turn them into **controlled components**.

With React **internal form components**, this can be done by using the **value** attribute. It supports an expression as the value that points to state.

```js
const [firstname, setFirstname] = useState("Alice");

return <input type="text" value={firstname} />;
```

In this example, the value of the firstname state is displayed in the input but this is only part of the solution. Right now, the firstname state isn't changed by the input. To accomplish that, the **onChange** event can be handled. It will fire on each keystroke.

```js
const [firstname, setFirstname] = useState("Alice");

return (
  <input
    type="text"
    value={firstname}
    onChange={(e) => setFirstname(e.target.value)}
  />
);
```

In the handler we can call the set function. We get the current input value from the event object passed into the handler function. **Event** object has a **target** property that is a reference to the DOM element object that caused the event to fire. From that, the value can be read, which contains a string that represents the input's contents.

If the input is a checkbox or radio, we use **checked** instead of **value**. This input is now a **controlled component** because the state of it is now fully controlled by React.

Now that the value of the input is controlled by React state, we can, of course benefit from all the goodness it brings, pass the state to other components or control the input's value outside of the input components using the set function for example.

## Forms

When the **input** is used on a **form**, the **onSubmit** event can be handled and the handler gets an event object on which **preventDefault** has to be called. It prevents browsers from doing the default submit action. May be you are used to returning **false** from the submit function to prevent the default action, but that's not supported in React.

```js
const [firstname, setFirstname] = useState("Alice");
const submit = (e) => {
  e.preventDefault();
  // submit first name to API
};
return (
  <form onSubmit={submit}>
    <input
      type="text"
      value={firstname}
      onChange={(e) => setFirstname(e.target.value)}
    />
  </form>
);
```

When multiple inputs are used, it is common practice to declare state as an object.

```js
const [person, setPerson] = useState({ firstname: "Alice", lastname: "Doe" });
const submit = (e) => {
  e.preventDefault();
  // submit person to API
};
return (
  <form onSubmit={submit}>
    <input
      type="text"
      value={person.firstname}
      onChange={(e) => setPerson({ ...person, firstname: e.target.value })}
    />
    <input
      type="text"
      value={person.lastname}
      onChange={(e) => setPerson({ ...person, lastname: e.target.value })}
    />
  </form>
);
```

Instead of giving each input it's own **onChange** handler, this is also a commonly used pattern:

```js
const [person, setPerson] = useState({ firstname: "Alice", lastname: "Doe" });
const submit = (e) => {
  e.preventDefault();
  // submit person to API
};
const change = (e) => setPerson({ ...person, [e.target.name]: e.target.value });
return (
  <form onSubmit={submit}>
    <input type="text" value={person.firstname} onChange={change} />
    <input type="text" value={person.lastname} onChange={change} />
  </form>
);
```

This syntax is known as **computed property name** in Javascript. `e.target.name` refers to the name attribute, which is now present on all inputs. The value of the attribute must correspond to the property name in the state object for this to work.

## Working with Input Components

Armed with above knowledge, We are going to see how a new feature around house bids work in the application.

In the **House** component there is now a section that shows bids on the house and there is an option to submit the new bid. This functionality is provided by a new component called **Bids**.

Let's add a new hook called **useBids**

```js
import { useEffect, useState } from "react";
import useGetRequest from "./useGetRequest";

const useBids = (houseId) => {
  const [bids, setBids] = useState([]);
  const { get, loadingState } = useGetRequest(`/api/bids/${houseId}`);

  useEffect(() => {
    const fetchBids = async () => {
      const bids = await get();
      setBids(bids);
    };
    fetchBids();
  }, [get]);
  const postBid = async (bid) => {
    await fetch(`/api/bids/${bid.houseId}`, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify(bid),
    });
  };

  const addBid = (bid) => {
    postBid(bid);
    setBids([...bids, bid]);
  };
  return { bids, loadingState, addBid };
};

export default useBids;
```

Now add Bids component as below:

```js
import useBids from "../hooks/useBids";

const { default: loadingStatus } = require("../helpers/loadingStatus");
const { useState } = require("react");
const { LoadingIndicator } = require("./loadingIndicator");
const { default: currencyFormatter } = require("../helpers/currencyFormatter");

const Bids = ({ house }) => {
  const { bids, loadingState, addBid } = useBids(house.id);

  const emptyBid = {
    houseId: house.id,
    bidder: "",
    amount: 0,
  };

  const [newBid, setNewBid] = useState(emptyBid);

  if (loadingState !== loadingStatus.loaded) {
    return <LoadingIndicator loadingState={loadingState} />;
  }

  const onBidSubmitClick = () => {
    addBid(newBid);
    setNewBid(emptyBid);
  };

  return (
    <>
      <div className="row mt-4">
        <div className="col-12">
          <table className="table table-sm">
            <thead>
              <tr>
                <th>Bidder</th>
                <th>Amount</th>
              </tr>
            </thead>
            <tbody>
              {bids.map((b) => (
                <tr key={b.id}>
                  <td>{b.bidder}</td>
                  <td>{currencyFormatter.format(b.amount)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
      <div className="row">
        <div className="col-5">
          <input
            id="bidder"
            className="h-100"
            type="text"
            value={newBid.bidder}
            onChange={(e) => setNewBid({ ...newBid, bidder: e.target.value })}
            placeholder="Bidder"
          />
        </div>
        <div className="col-5">
          <input
            id="amount"
            className="h-100"
            type="number"
            value={newBid.amount}
            onChange={(e) =>
              setNewBid({ ...newBid, amount: parseInt(e.target.value) })
            }
            placeholder="Amount"
          />
        </div>
        <div className="col-2">
          <button className="btn btn-primary" onClick={onBidSubmitClick}>
            Add
          </button>
        </div>
      </div>
    </>
  );
};

export default Bids;
```

Now add this **Bids** component in **House** component.

```js
import defaultPhoto from "../helpers/defaultPhoto";
import React, { useContext } from "react";
import { navigationContext } from "./app";
import Bids from "./bids";
import currencyFormatter from "../helpers/currencyFormatter";

const House = () => {
  const { param: house } = useContext(navigationContext);
  return (
    <div className="row">
      <div className="col-6">
        <div className="row">
          <img
            className="img-fluid"
            src={
              house.photo ? `./houseImages/${house.photo}.jpeg` : defaultPhoto
            }
            alt="House pic"
          />
        </div>
      </div>
      <div className="col-6">
        <div className="row mt-2">
          <h5 className="col-12">{house.country}</h5>
        </div>
        <div className="row">
          <h3 className="col-12">{house.address}</h3>
        </div>
        <div className="row">
          <h2 className="themeFontColor col-12">
            {currencyFormatter.format(house.price)}
          </h2>
        </div>
        <div className="row">
          <div className="col-12 mt-3">{house.description}</div>
        </div>
        <Bids house={house} />
      </div>
    </div>
  );
};

export default House;
```

## Non-input Form Components

There are two input components that work differently than their HTML equivalent in React.

- **textarea**

  - In HTML text is placed as the innerHTML. Because of this making this a **controlled component** while maintaining the same usage is difficult.
  - That's why the React version works similarly to an input component in React. We have the value attribute that can point to state and onChange event with a handler that changes the state.

  HTML

  ```html
  <textarea>
    Some text
  </textarea>
  ```

  React

  ```js
  <textarea value={state} onChange={change} />
  ```

- **select**

  - This faciliates drop-down list. In HTML we determine which option is selected using the **selected** attribute.
  - This makes it hard to work with state, and that's why in react we use it again like input. **state** below determines which item is selected. If state contains the string "option1" the first one will be selected. And when user select an option, **change** will fire, which again like input can change the state using the value property on the target object available on the event object.

  HTML

  ```html
  <select>
    <option value="option1">1</option>
    <option selected value="option2">2</option>
  </select>
  ```

  REACT

  ```js
  <select value={state} onChange={change}>
    <option value="option1">1</option>
    <option value="option2">2</option>
  </select>
  ```

## Uncontrolled Components

Normally using **controlled component** is definitely the way to go. However, there can be a lot of work because we need to write an event handler for every way your data can change and pipe all of the input state through a React component.
This can become annoying when you are converting a pre-existing codebase to React, for example, or if you want a quick and dirty solution. In these cases, you might consider to use uncontrolled components, may be temporarily.
Having said that, in software development temporary things become permanent. So it's worth investing extra time to do it the controlled way from the start.

Anyways, here is how uncontrolled components work.

```js
const Form = () => {
  const inputEl = useRef(null);
  const submit = (e) => {
    e.preventDefault();
    const inputValue = inputEl.current.value;
    // process input value
  };
  return (
    <form>
      <input ref={inputEl} type="text" />
      <input type="submit" value="Submit" />
    </form>
  );
};
```

Components can have **ref** attribute that points to the resulting variable. The JavaScript DOM object representing the current element in the browser becomes available in the variable. When it's time to submit the form value of each input can be read by using the **current** property on the refs, which contains the object.

Note we are not using react's state mechanism at all here. Instead the native state of the browser elements is used. React doesn't have any control over that, hence the name **uncontrolled components**.

It's also possible to use an expression for the **defaultValue** attribute of a component that uses some state or a variable. This value will be set on first render. After that the input will be uncontrolled.

```jsx
<input ref="{inputEl}" type="text" defaultValue={val} />
```

A file input is always an uncontrolled component because it's value can only be set by a user, not in code. That means react doesn't have a way to put a state value into it.

```js
const Form = () => {
  const inputEl = useRef(null);
  const submit = (e) => {
    e.preventDefault();
    const selectedFile = inputEl.current.files[0].name;
    // process input value
  };
  return (
    <form>
      <input ref={inputEl} type="file" />
      <input type="submit" value="Submit" />
    </form>
  );
};
```

So when using input type as **file**, you will always have to use a ref. In the submit function, the files can be read from the element object.

## Form Libraries

If you need more out of the box functionalities such as:

- Validation
- Error messages
- Handling form submissions
- Making state handling easier

You could take a look at an external library that helps with forms. The most popular one is called **[Formik](http://formik.org)**. Keep in mind that it's built on same principles of controlled components and managing state.

# Application Design

Application design is something you need when starting with a new project, but also during the lifetime of the project when new features are added. In both case Mocks are probably made, which basically is a sketch depicting how the UI should look like. And now you are asked to implement it. The first thing to do is to determine the components that have to be built with the single responsibility principle in mind that basically tells us that every component should do just one thing.
In this case, since this is the start of a new application, there first is the component that contains all the other components, the **root component**.

Being the container for all other components is its responsiblity. In the root component is the **Banner**, the **HouseList** and the **HouseRow**.

You don't want components that do too much because they tend to become unmaintainable. But keep in mind that design is a living thing.

You shouldn't go too far, but also be sure you go far enough when creating the component hierarchy.

In some point in the course, we added an **Add** button, for example, and we added it as part of **HouseList**. It's good idea to create a separate component for it because this is a typical example of a single responsibility violation.

A component that lists something shouldn't also be adding something. These are two separate tasks.
In the **House** composition, there's similar example. The picture and information about the house is perfectly fine to have in the **House** component, and the **House** component renders a **Bids** component. Rightnow we have a list of Bids and the adding of Bids in one component. But now that we take a look at it from a design perspective, that might not have been such a good idea.

Let's go through all the design steps first. After applying the Bids fix and identifying all the other components, a hierarchy like below can be created.

- App
  - Banner
  - HouseList
    - HouseRow
  - House
    - BidList
    - AddBid

With this hierarchy in hand we can go to the next step.

## File Structure & Building a Static Version

In demo application we have used a system that groups the files by type.

- components
  - app.js
  - houseList.js
  - house.js
  - ...
- helpers
  - currencyFormatter.js
  - loadingStatus.js
  - navValues.js
  - ...
- hooks
  - useHouses.js
  - useBids.js
  - useGetRequest.js
  - ...

This works great for now, but when this application grows beyond supporting just houses, the number of files and components, for instance, will grow to a point where it's getting hard to find a particular file on the list. In that case, it would be an option to create a directory per feature and put all the files for the feature directly under that and have a common directory for everything that is needed by each feature while having the root component at the root directory level.

- app.js
- houses
  - houseList.js
  - house.js
  - useHouses.js
  - useBids.js
  - ...
- customers
  - customerList.js
  - useCustomers.js
  - ...
- common
  - currencyFormatter.js
  - loadingStatus.js
  - navValues.js
  - useGetRequest.js
  - ...

Using this structure, you could consider to rename the component that you consider the main component for a feature to `index.js`. The advantage is that this component can now be imported by just using the directory name omitting the file name. This is a **webpack** feature, so it only worksif you're using that. `Next.js` and `Create React App` use it by default.

- app.js
- houses
  - index.js : `import HouseList from "./houses"`
  - house.js
  - useHouses.js
  - useBids.js
  - ...
- customers
  - index.js : `import CustomerList from "./customers"`
  - useCustomers.js
  - ...
- common
  - currencyFormatter.js
  - loadingStatus.js
  - navValues.js
  - useGetRequest.js
  - ...

However drawback is, there would be multiple files showing up in the tab bar of the code editor that are named `index.js` and it's hard to tell the difference.

React doesn't have an opinion on how you should put files into directories, so you can do it however you think is right for your application. But limiting yourself to a directory nesting level of three or four is a good idea to avoid getting overwhelmed by the sheer number of directories.

## Create a First Application Version

It's time to open a code editor and make a first version of the app, a version that doesn't support interactivity yet, but just renders the UI. In other words, you should create it without the use of state. Just use props to pass on data in the hierarchy. Creating a one-way data flow. And as a data source for now use something in-memory. Only when you have that working, state can be introduced.

## Introducing State

Creating state is easy enough, but it should be in the application as little as possible to avoid re-renders. To identify if data should be state, this checklist can be used:

- Data is probably not state for a component when it's passed in by a parent using props.
- The data involved should also change over time. If it's static data then there is no point putting that in state.
- When the data can be computed, it probably isn't state either.

So what about houses array, is that state?
For initial app it won't be passed in as a prop. We start out with an empty array, and then it changes after loading and we can't compute it. So that qualifies a state.

And if we were to include a total number of items counter in our **HouseList** component, well that can easily be determined by reading the length property of the data array, so no need to use state. Computing it is trivial.

What about the default house photo?
It's a static string so no.

The Bids are because of the same reasons as houses.

And the NewBid, is that state?
Yes, can't be computed, it's changed by the user, and it is not passed in via a prop.

## Placing and Applying state

We have learned that state can live throughout the application because it can be passed around using props or contacts, for example. But there is only one component that owns a particular piece of state. That's the place where the state is introduced. Place state as low as possible in the component hierarchy.

Lets take example of NewBid state. What component use that state?

- App
  - Banner
  - HouseList
    - HouseRow
  - House
    - BidList
    - AddBid

AddBid is the only one, so that should be the owner. Now the Bid state, which component use it?
**BidList** obviously, but also **AddBid**. Because they both use it, we can walk up the tree to look for a common parent. In this case there is **House**, so the state should be there.

If a common parent isn't so easily found, you can also consider creating one, and the sole purpose of this component could be to introduce state. Now in future there may be a requirement that **HouseList** should should show highest Bid in the table. Now there are three component the use the state. Making **House** the state owner doesn't work anymore. We should look higher up in the component tree to see where it can go. In this case, it can be introduced in the **App** root component because that is the common parent for all components that use it.

This is called lifting state up because the new location is higher up in the component hierarchy.

How to apply state?
We could just use the state hook i.e. **useState** in the designated component, or we could create a custom hook that hold the state. When using a custom hook, the component that uses it is still the owner. All state and effects inside the hook are considered part of the component.

Using **context** is another option if you don't want to pass the state via props to many components or if you want to implement a navigation technique that is basically state available for the whole application, global state.

## Adding Inverse Data Flow

The normal data flow in React is from higher-level components in the hierarchy to lower-level ones because we are passing on data via props. This is often referred to as one way data flow of React.
But sometimes a child component wants something done by it's parent. The most common thing it wants is to change the state of the parent. We do that by invering the data flow. We've already seen that in action when we passed functions to children. That's React's way to inverse flow.

When children call such a function, you could say that they are swimming against the one-way stream. In the design phase it's good idea to identify the places where that should happen.

## Implementing Design Changes

While going through the all design steps for the whole application, we discovered that the functionality to list bids and to add them should be separated. We also saw that the two components should be siblings with the **House** component as the parent and that on that component level the bid's state should live. And there is one other piece of state for the new bid that should be in the **AddBid** component.

Now let's implement the design changes.
