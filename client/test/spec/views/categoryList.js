describe('Testing CategoryList View', function () {

  // setup
  beforeEach(module('client'));

  var TestCtrl, $rootScope, $compile, createController, view, $scope;
  beforeEach(inject(function($controller, $templateCache, _$rootScope_, _$compile_, _$httpBackend_) {
    $rootScope = _$rootScope_;
    $scope = $rootScope.$new();
    $compile = _$compile_;
    $httpBackend = _$httpBackend_
    $httpBackend.whenGET(/^.*/).respond('');

    createController = function(data) {
      var html = '<div>' + $templateCache.get('templates/categoryList.html') + '</div>';
      TestCtrl = $controller('categoryListCtl', { $scope: $scope, $rootScope: $rootScope });
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

});