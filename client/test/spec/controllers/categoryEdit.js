describe('categoryEdit tests', function() {
  beforeEach(module('client'))
 	var listCtl, scope, list, service, q, deferred, state, nextState, dummy;

	beforeEach(function(){
		service = {
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

		state = {
      params: { 
        id: 39
      }, 
			go: function(s, p) {
				nextState = {name:s, params: p};
			}
		};

    dummy = {
      id: 39
    }
	});


	beforeEach(inject(function($controller, $rootScope, $q){
		q = $q;
		scope = $rootScope.$new();
		listCtl = $controller('categoryEditCtl', {
			$scope: scope,
			categoryService: service,
			$state: state
		})
	}))

  
	it('should get category during init', function() {  	
		spyOn(service, 'Get').and.callThrough();
		scope.init();
		expect(service.Get).toHaveBeenCalled();
	});


  it('should have an update method that calls the update service', function() {    
    spyOn(service, 'Update').and.callThrough();
    scope.init();
    scope.update(dummy);
    expect(service.Update).toHaveBeenCalled();
  });

  it('after the update method, the controller should redirect to the list', function() {    
    scope.init();
    scope.update(dummy);
    deferred.resolve('ok');
    scope.$root.$digest();
    expect(nextState.name).toBe('category');
  });

});