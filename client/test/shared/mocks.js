mockGenericService = function(){
    window.service = {
          List: function () {
            deferred = q.defer();
              return deferred.promise;
          },
          Get: function (id) {
            deferred = q.defer();
              return deferred.promise;
          },
          Update: function (category) {
            deferred = q.defer();
              return deferred.promise;
          },
          Create: function (category) {
            deferred = q.defer();
              return deferred.promise;
          }
    }
}

mockState = function() {
    window.state = {
      params: stateParams, 
      go: function(s, p) {
        nextState = {name:s, params: p};
      }
    }

    window.nextState = ''
}

test_init = function() {
  window.listCtl, window.scope, window.list, window.service, window.q, window.deferred, window.state, window.nextState,
  window.dummy; window.stateParams = {};

  beforeEach(module('client'))
}

