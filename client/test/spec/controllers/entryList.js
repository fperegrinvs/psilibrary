describe('entryList tests', function() { 	
  test_init()

  // mock service
  beforeEach(mockState);
  beforeEach(mockGenericService);

  beforeEach(inject(function($controller, $rootScope, $q, $route, $state){
    q = $q;
    realState = $state
    route = $route
    scope = $rootScope.$new();
    listCtl = $controller('entryListCtl', {
      $scope: scope,
      entryService: service,
      $state: state
    })
  }))

  it('should get list during init', function() {    
    spyOn(service, 'List').and.callThrough();
    scope.init();
    expect(service.List).toHaveBeenCalled();
  });

  it('should store list result on scope', function() {
    scope.init();
    deferred.resolve({id:'oi'});
    scope.$root.$digest();
    expect(scope.data.id).toBe('oi');
  });

  it('should store error message if listing fails', function() {
    scope.init();
    deferred.reject('Error');
    scope.$root.$digest();
    expect(scope.msg).toEqual({error:'Error'});
  });

  it('should have a create method that redirect to create page', function(){
    scope.create()
    expect(nextState.name).toBe('entryCreate');
  });

  it('should have a edit method that redited to edit page', function(){
    scope.edit(3)
    expect(nextState.name).toBe('entryEdit');
  });

  it('should edit method should require the entry id', function(){
    scope.edit()
    expect(nextState.name).not.toBe('entryEdit');
  });

  it('should include error message if edit have no parameter', function(){
    scope.edit()
    expect(scope.msg.error).not.toBe(undefined);
  });

  it('should store error message if and error is passed on state params', function(){
    state.params = {error: 'An error'};
    scope.init()
    expect(scope.msg).toEqual(state.params);
 })


  it('check if list route exist', function(){
    r = realState.get('entry')
    expect(r).not.toBe(null)
  })

});