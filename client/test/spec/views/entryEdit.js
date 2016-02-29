describe('Testing EntryEdit View', function () {

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
      var html = '<div>' + $templateCache.get('templates/entryEdit.html') + '</div>';
      TestCtrl = $controller('entryCreateCtl', { $scope: $scope, $rootScope: $rootScope });
      $scope.data = data;
      view = $compile(angular.element(html))($scope);
      $scope.$digest();
      return $scope;
    };
  }));

  // tests
  it('Checking the id field, with data', function() {
    createController({id: 3});
    expect(view.find("#id").val()).toEqual('3')
  });

  it('id field should be ommited if no id on data', function() {
    createController({name: 3});
    expect(view.find("#id").length).toEqual(0)
  });

  it('should have a save button', function(){
    createController({name: 3});
    expect(view.find(".save").length).toEqual(1)
  })

  it('save button should call the save method', function(){
    createController({name: 3});
    spyOn($scope, 'save');
    $(view).find('.save')[0].click();
    expect($scope.save).toHaveBeenCalled();
  })

  it('Checking the title field', function() {
    createController({title: 'test'});
    expect($scope.dataForm.title.$modelValue).toEqual('test')
  });

  it('Title field should be required', function() {
    createController({});
    $scope.dataForm.$setSubmitted();
    expect($scope.dataForm.title.$error.required).not.toBe(undefined)
  })
});