describe('Testing EntryList View', function () {

  // setup
  beforeEach(module('client'));

  var TestCtrl, $rootScope, $compile, createController, $scope;
  beforeEach(inject(function($controller, $templateCache, _$rootScope_, _$compile_, _$httpBackend_) {
    $rootScope = _$rootScope_;
    $scope = $rootScope.$new();
    $compile = _$compile_; 
    $httpBackend = _$httpBackend_
    $httpBackend.whenGET(/^.*/).respond('');

    createController = function(data) {
      var html = '<div>' + $templateCache.get('templates/entryList.html') + '</div>';
      TestCtrl = $controller('entryListCtl', { $scope: $scope, $rootScope: $rootScope });
      $scope.data = data;
      view = $compile(angular.element(html))($scope);
      $scope.$digest();
      return $scope;
    };
  }));

  // tests
  it('Checking the edit buttons', function() {
    createController([{id: 1},{id: 2},{id: 3},{id: 4}]);
    expect(view.find(".edit").length).toEqual(4)
  });

  it('Checking the edit click', function() {
    var scope = createController([{id: 1},{id: 2},{id: 3},{id: 4}]);
    spyOn(scope, 'edit');
    $(view).find('.edit')[3].click();

    expect(scope.edit).toHaveBeenCalled();
  });

  it('Checking the create button', function() {
    createController([{id: 1},{id: 2},{id: 3},{id: 4}]);
    expect(view.find(".create").length).toEqual(1)
  });

  it('Checking the create click', function() {
    var scope = createController([]);
    spyOn(scope, 'create');

    $(view).find('.create')[0].click();

    expect(scope.create).toHaveBeenCalled();
  });

  it('first column should be title', function(){
    var scope = createController([{id: 1, title: 'Test'}]);
    var column = $(getColumn(0)).text(); 
    
    expect(column).toBe('Test');
  })

  it('2nd column should be abstract', function(){
    var scope = createController([{id: 1, abstract: 'Test'}]);
    var column = $(getColumn(1)).text(); 
    
    expect(column).toBe('Test');
  })

  it('if abstract is too big, it should show only the first 50 letters, followed by ...', function(){
    var scope = createController([{id: 1, abstract: 'Lorem ipsum dolor sit amet, consectetur adipiscing volutpat.'}]);
    var column = $(getColumn(1)).text(); 
    
    expect(column).toEqual('Lorem ipsum dolor sit amet, consectetur adipiscing...');
  })

  it('3rd column should be author', function(){
    var scope = createController([{id: 1, author: 'Joe'}]);
    var column = $(getColumn(2)).text(); 
    
    expect(column).toBe('Joe');
  })
});
