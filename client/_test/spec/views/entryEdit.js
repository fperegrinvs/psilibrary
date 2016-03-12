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

  it('should have a input with a list of categories', function(){
    createController({id: 3}, categories);
    expect(view.find(".categoriesList").children().length).toEqual(3);
  })

  it('should have a list of entry`s categories', function(){
    createController({id: 3, categories: categories}, categories);
    expect(view.find(".category").length).toEqual(2);
  });

  it('it should have a button to remove a category', function(){
    createController({id: 3, categories: categories}, categories);
    expect(view.find(".remove-category").length).toEqual(2);
  });

  it('the remove button should call the removeCategory method', function(){
    createController({id: 3, categories: categories}, categories);
    spyOn($scope, 'removeCategory');
    $(view).find('.remove-category')[0].click();
    expect($scope.removeCategory).toHaveBeenCalledWith(categories[0]);
  });

});