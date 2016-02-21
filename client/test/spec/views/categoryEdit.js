describe('Testing CategoryEdit View', function () {

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
      var html = '<div>' + $templateCache.get('templates/categoryEdit.html') + '</div>';
      TestCtrl = $controller('categoryCreateCtl', { $scope: $scope, $rootScope: $rootScope });
      $scope.data = data;
      view = $compile(angular.element(html))($scope);
      $scope.$digest();
      return $scope;
    };
  }));

  // tests
  it('Checking the id field, no data', function() {
    createController({});
    expect(view.find("#id").length).toEqual(0)
  });

  it('Checking the id field, with data', function() {
    createController({id: 3});
    expect(view.find("#id").length).toEqual(1)
    expect(view.find("#id").val()).toEqual('3')
  });

 it('Checking the name field', function() {
    createController({name: 'test'});
    expect(view.find("#name").length).toEqual(1)
    expect(view.find("#name").val()).toEqual('test')
  });

  it('Checking the save button', function() {
    createController({});
    expect(view.find(".save").length).toEqual(1)
  });

  it('Checking the save click', function() {
    var scope = createController({id:1, name:'teste'});
    spyOn(scope, 'update');

    $(view).find('.save')[0].click();

    expect(scope.update).toHaveBeenCalled();
  });

});