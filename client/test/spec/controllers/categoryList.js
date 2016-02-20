describe('categoryList tests', function() {
  beforeEach(module('client'))
 	var listCtl, scope, list, service, q, deferred, state, nextState;

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
			go: function(s, p) {
				nextState = {name:s, params: p};
			}
		}
	});


	beforeEach(inject(function($controller, $rootScope, $q){
		q = $q;
		scope = $rootScope.$new();
		listCtl = $controller('categoryListCtl', {
			$scope: scope,
			categoryService: service,
			$state: state
		})
	}))

  
	it('should get list during init', function() {  	
		spyOn(service, 'List').and.callThrough();
		scope.init();
		expect(service.List).toHaveBeenCalled();
	});

  it('it should store the list on scope', function(){
    scope.init();
    deferred.resolve({test: 'oi'});
    scope.$root.$digest();
    expect(scope.data.test).toBe('oi')
  })

  it('should have a create method that redirects to the create page', function() {    
    scope.init();
    scope.create();
    expect(nextState.name).toBe('categoryCreate');
  });

  it('should have an edit method that redirects to the edit page', function() {    
    scope.init();
    scope.edit(1);
    expect(nextState.name).toBe('categoryEdit');
    expect(nextState.params.id).toBe(1);
  });

});