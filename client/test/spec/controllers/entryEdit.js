describe('entryEdit tests', function() { 	
  test_init()

  beforeEach(mockGenericService);
  beforeEach(mockState);

  //global variables
  beforeEach(function(){
    dummy = {
      id: 39,
      name: 'test'
    }

    stateParams = {id: 39}
  });

  beforeEach(inject(function($controller, $rootScope, $q){
    q = $q;
    scope = $rootScope.$new();
    listCtl = $controller('entryEditCtl', {
      $scope: scope,
      entryService: service,
      $state: state
    })
  }))

  it('should call service.get on init', function(){
  	spyOn(service, 'Get').and.callThrough();
  	scope.init();
    expect(service.Get).toHaveBeenCalled();
  });

  it('the call of service.get should have the id param of current state', function(){
  	spyOn(service, 'Get').and.callThrough();
  	scope.init();
  	expect(service.Get).toHaveBeenCalledWith(39)
  });

  it('if current state lacks id, redirect to list state with error message', function(){
	state.params = {}
	scope.init();
    expect(nextState.name).toBe('entryList');
    expect(nextState.params.error).not.toBe(undefined);
  });

  it('should call service update on save', function(){
  	spyOn(service, 'Update').and.callThrough();
  	scope.save();
    expect(service.Update).toHaveBeenCalled();
  });

  it('on save, should pass scope data as param to update service', function(){
  	spyOn(service, 'Update').and.callThrough();
  	scope.data = {id: '1', 'name': 'test'}
  	scope.save();
  	expect(service.Update).toHaveBeenCalledWith(scope.data)  	
  });

  it('after save, should store success message if ok', function(){
  	scope.save();
  	deferred.resolve('ok');
    scope.$root.$digest();
  	expect(scope.msg.success).not.toBe(undefined);
  })

  it('after save, should store error message if error', function(){
  	scope.save();
  	deferred.reject('error');
    scope.$root.$digest();
  	expect(scope.msg.error).toBe('error');
  })

});