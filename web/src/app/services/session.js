/*
Copyright 2015 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

var { browserHistory, createMemoryHistory } = require('react-router');

const AUTH_KEY_DATA = 'authData';

var _history = createMemoryHistory();

var session = {

  init(history=browserHistory){
    _history = history;
  },

  getHistory(){
    return _history;
  },

  setUserData(userData){
    localStorage.setItem(AUTH_KEY_DATA, JSON.stringify(userData));
  },

  getUserData(){
    var hiddenDiv = document.getElementById("bearer_token");
    if(hiddenDiv!== null){
      let content = JSON.parse(hiddenDiv.textContent);
      hiddenDiv.remove();
      this.setUserData(content.session);
    }

    var item = localStorage.getItem(AUTH_KEY_DATA);
    if(item){
      return JSON.parse(item);
    }

    return {};
  },

  clear(){
    localStorage.clear()
  }

}

module.exports = session;
