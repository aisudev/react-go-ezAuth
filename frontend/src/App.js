import React from "react";
import SignInScreen from "./screens/SignInScreen";
import SignUpScreen from "./screens/SignUpScreen";
import { BrowserRouter as Router, Switch, Route, Redirect } from 'react-router-dom'
import HomeScreen from "./screens/HomeScreen";
import Controller, { useController } from './controllers/Controller'
import withController from './hoc/withController'

function App() {

  return (
    <Router>
      <Switch>

        <AuthRoute path='/signin' component={SignInScreen} />
        <AuthRoute path='/signup' component={SignUpScreen} />
        <ProtectRoute path='/' component={HomeScreen} />

      </Switch>
    </Router>
  );
}

function AuthRoute({ component, path }) {
  const controller = useController()

  return !controller.user ? (<Route component={component} path={path} />) : (<Redirect to='/' />)
}

function ProtectRoute({ component, path }) {
  const controller = useController()

  return controller.user ? (<Route component={component} path={path} />) : (<Redirect to='/signin' />)
}

export default withController(App)(Controller)
