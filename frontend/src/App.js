import React from "react";
import SignInScreen from "./screens/SignInScreen";
import SignUpScreen from "./screens/SignUpScreen";
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import HomeScreen from "./screens/HomeScreen";
import Controller from './controllers/Controller'
import withController from './hoc/withController'

function App() {

  return (
    <Router>
      <Switch>
        <Route path='/signin' component={SignInScreen} />
        <Route path='/signup' component={SignUpScreen} />
        <Route path='/' component={HomeScreen} />
      </Switch>
    </Router>
  );
}

export default withController(App)(Controller)
