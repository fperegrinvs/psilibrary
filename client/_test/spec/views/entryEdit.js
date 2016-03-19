describe('Testing EntryEdit View', function () {

  // setup
  beforeEach(module('client'));
  beforeEach(mockGenericService);

  var TestCtrl, $rootScope, $compile, createController, $scope;
  var categories = [{id: 1, name: 'oi'}, {id: 2, name: 'teste'}];
  beforeEach(inject(function($controller, $templateCache, _$rootScope_, _$compile_, _$httpBackend_) {
    $rootScope = _$rootScope_;
    $scope = $rootScope.$new();
    $compile = _$compile_;
    $httpBackend = _$httpBackend_
    $httpBackend.whenGET(/^.*/).respond('');

    createController = function(data, categories) {
      var html = '<div>' + $templateCache.get('app/templates/entryEdit.html') + '</div>';
      TestCtrl = $controller('entryCreateCtl', { $scope: $scope, $rootScope: $rootScope });
      $scope.data = data;

      if (categories) {
        $scope.categories = categories;
      }

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
  });

  it('Checking the author field', function() {
    createController({author: 'test'});
    expect($scope.dataForm.author.$modelValue).toEqual('test')
  });

  it('Author field should be required', function() {
    createController({});
    $scope.dataForm.$setSubmitted();
    expect($scope.dataForm.author.$error.required).not.toBe(undefined)
  });

  it('should have title "Novo registro" if creating a new entry', function(){
    createController({});
    expect(view.find(".title").text()).toEqual('Novo registro')
  });

  it('should have title "Editando registro" when editing an existing entry', function(){
    createController({id: 3});
    expect(view.find(".title").text()).toEqual('Editando registro')
  });

  it('should have an input for publishdate', function(){
    date = new Date()
    createController({publishDate: date});
    expect($scope.dataForm.publishDate.$modelValue).toEqual(date)
  })

// usando plugin bootstrap
  it('should have a input with a list of categories', function(){
    createController({id: 3}, categories);
    expect(view.find("#bootstrap-duallistbox-nonselected-list_").children().length).toEqual(2);
  })

  it('should have a list of entry`s categories', function(){
    createController({id: 3, categories: categories}, categories);
    expect(view.find("#bootstrap-duallistbox-selected-list_").children().length).toEqual(2);
  });

  it('should have a input with entry abstract', function(){
    createController({abstract: 'hello world'});
    expect($scope.dataForm.abstract.$modelValue).toEqual('hello world')
  });

  it('shoud have a input with entry content', function(){
   createController({content: 'hello world'});
    expect($scope.dataForm.content.$modelValue).toEqual('hello world')
  });

  it('shoud have a input with journal', function(){
   createController({journal: 'hello world'});
    expect($scope.dataForm.journal.$modelValue).toEqual('hello world')
  });

  it('should have a selection box to select the entry type', function(){
    var entryType = {id: 1, name: 'teste'};
    createController({entryType: entryType});
    expect($scope.dataForm.entryType.$modelValue).toEqual(entryType)
  });

  it('entry type should be required', function(){
    createController({});
    $scope.dataForm.$setSubmitted();
    expect($scope.dataForm.entryType.$error.required).not.toBe(undefined)
  })

  it('abstract should be required', function(){
    createController({});
    $scope.dataForm.$setSubmitted();
    expect($scope.dataForm.abstract.$error.required).not.toBe(undefined)
  })

});