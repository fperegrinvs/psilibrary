describe('entryCreate tests', function() { 	
  test_init()

  window.categories = [{id:1, name:'cat1'},{id:2, name:'cat2'}]
  // mock service
  beforeEach(mockState);
  beforeEach(mockGenericService);
  categoryService = {
    List: function() {return window.categories;}
  }

  beforeEach(inject(function($controller, $rootScope, $q, $state){
    categoryService = jQuery.extend(true, {}, service);
    q = $q;
    realState = $state
    scope = $rootScope.$new();
    listCtl = $controller('entryCreateCtl', {
      $scope: scope,
      entryService: service,
      $state: state,
      categoryService: categoryService,
    })
  }))

  it('should have an init object after init', function(){
  	scope.init();
  	expect(scope.data).toBe(undefined)
  })

  it('should have a save method that calls service create', function(){
  	spyOn(service, 'Create').and.callThrough();
    scope.data = {title: 'Oi', author: 'Hello'}
  	scope.save();
    expect(service.Create).toHaveBeenCalled();
  })

  it('on save, service creat should receive scope data', function(){
	spyOn(service, 'Create').and.callThrough();
  	scope.data = dummy
  	scope.save();
  	expect(service.Create).toHaveBeenCalledWith(dummy)
  })

  it('should display a success message on service success', function(){
  	scope.save();
    deferred.resolve('ok');
    scope.$root.$digest();
    expect(scope.msg.success).not.toBe(undefined);
    expect(scope.msg.error).toBe(undefined);
  });

  it('should store error message if create fails', function(){
  	scope.save();
    deferred.reject('Error');
    scope.$root.$digest();
    expect(scope.msg).toEqual({error: 'Error'});
  });

  it('check if edit rounte exists', function() {
    var r = realState.get('entryCreate');
    expect(r).not.toBe(null);
  })

  it('should load a categories list', function(){
    scope.init();
    deferred.resolve(window.categories);
    scope.$root.$digest();
    expect(scope.categories).toEqual(window.categories);
  })

});