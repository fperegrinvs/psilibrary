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
  window.dummy; window.stateParams = {}; window.error; window.$httpBackend; window.route; window.view;

  beforeEach(module('client'))
}

  getResult = function(data){
    dummy = data;
  }

  getError = function(e){
      error = e;
  }

  serviceOk = function(method, param, url, response, body){
      var call = !body ? $httpBackend.expectGET(url) : $httpBackend.expectPOST(url, body);
      call.respond(200, response);

      method(param).then(getResult, getError);

      $httpBackend.flush();
      expect(dummy).not.toBe(null);
      expect(error).toBe(null)
    }

  serviceFail = function(method, param, url, response, body){
      var call = !body ? $httpBackend.expectGET(url) : $httpBackend.expectPOST(url, body);
      call.respond(500, response);

      method(param).then(getResult, getError);

      $httpBackend.flush();
      expect(dummy).toBe(null);
      expect(error).not.toBe(undefined)
    }

getColumn = function(column, row) {
  if (!row) {
    row = 0;
  }
  
  return $($(view).find('#table-data tbody tr')[row]).find('td')[column]
}

triggerValidation = function() {
    angular.element($('form')).scope().dataForm.$setSubmitted();
}
