describe('categoryCreate tests', function() {
  test_init()

	beforeEach(mockGenericService); 
  beforeEach(function(){
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
		listCtl = $controller('categoryCreateCtl', {
			$scope: scope,
			categoryService: service,
			$state: state
		})
	}))


  it('should have an empty object after init', function() {    
    scope.init();
    expect(scope.data).not.toBe(undefined);
  });


  it('should have an update method that calls the create service', function() {    
    spyOn(service, 'Create').and.callThrough();
    scope.init();
    scope.update(dummy);
    expect(service.Create).toHaveBeenCalled();
  });

  it('after the update method, the controller should redirect to the list', function() {    
    scope.init();
    scope.update(dummy);
    deferred.resolve('ok');
    scope.$root.$digest();
    expect(nextState.name).toBe('category');
  });

});