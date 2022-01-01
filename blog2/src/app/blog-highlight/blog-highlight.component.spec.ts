import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BlogHighlightComponent } from './blog-highlight.component';

describe('BlogHighlightComponent', () => {
  let component: BlogHighlightComponent;
  let fixture: ComponentFixture<BlogHighlightComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BlogHighlightComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BlogHighlightComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
