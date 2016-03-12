describe('categoryList tests', function() {
  test_init();

  beforeEach(mockState);
  beforeEach(mockGenericService);


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

  it('should store error message if and error is passed on state params', function(){
    state.params = {error: 'An error'};
    scope.init()
    expect(scope.msg).toEqual(state.params);
 })

});